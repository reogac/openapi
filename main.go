package main

import (
	"bytes"
	"fmt"
	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/orderedmap"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{})
}

type DataModel struct {
	id         string
	goType     string
	isArray    bool
	isMap      bool
	properties map[string]Property
	schema     *base.Schema
	enum       *Enum
}

func (m *DataModel) getGoType() string {
	if m.enum != nil {
		return m.id
	}
	return m.goType
}

type Enum struct {
	enumType string
	values   []string
}

type Property struct {
	id       string
	required bool
	m        *DataModel
}

func (p *Property) hasEmptyValue() bool {
	if p.m.isArray {
		return true
	}

	if p.m.isMap {
		return true
	}

	if p.m.goType == "string" {
		return true
	}

	return false
}

func (p *Property) writeTag() string {
	if p.required {
		return fmt.Sprintf("`json:\"%s\"`", p.id)
	} else {
		return fmt.Sprintf("`json:\"%s,omitempty\"`", p.id)
	}
}

type RequestBody struct {
	required bool
	content  DataModel
	desc     string
}

type Parameter struct {
	id       string
	in       string
	desc     string
	required bool
	m        DataModel
}

func (p *Parameter) hasEmptyValue() bool {
	return p.m.isArray || p.m.isMap || p.m.goType == "string"
}

func (p *Parameter) isPrimitive() bool {
	return isPrimitive(p.m.goType)
}

func (p *Parameter) getName(capitalize bool) string {
	if len(p.id) == 0 {
		return ""
	}

	parts := strings.FieldsFunc(p.id, func(c rune) bool {
		return c == ' ' || c == '-' || c == '_'
	})
	out := ""
	for i, part := range parts {
		if i == 0 {
			if part[0] == '5' {
				out = "five" + part[1:]
			} else if part[0] == '3' {
				out = "three" + part[1:]
			} else {
				out = strings.ToLower(string(part[0])) + part[1:]
			}
			if capitalize {
				out = strings.Title(out)
			}
		} else {
			out = out + strings.Title(part)
		}
	}
	return out
}

func (p *Parameter) getTypeDefinition() (def string) {
	if p.m.isArray {
		def = "[]" + p.m.goType
	} else if p.m.isMap {
		def = "map[string]" + p.m.goType
	} else {
		if p.required {
			if p.isPrimitive() {
				def = p.m.goType
			} else { //struct type
				def = "*" + p.m.goType
			}
		} else {
			if p.m.goType == "string" {
				def = p.m.goType
			} else { //other primitive types or struct
				def = "*" + p.m.goType
			}
		}
	}
	return
}
func (p *Parameter) isStringType() bool {
	if p.m.isArray || p.m.isMap {
		return false
	}
	return p.m.goType == "string"
}

func (p *Parameter) TypeName() string {
	goType := strings.Title(p.m.goType)
	if p.m.isArray {
		return "ArrayOf" + goType
	} else if p.m.isMap {
		return "MapOf" + goType
	}
	return goType
}

func (p *Parameter) getDefinition(capitalize bool) string {
	return fmt.Sprintf("%s %s", p.getName(capitalize), p.getTypeDefinition())
}

//write code for checking nil value of a parameter and add to the request
func (p *Parameter) writeParamCheck(prefix string) string {
	lines := []string{}

	var varStr string
	inStruct := len(prefix) > 0

	if inStruct {
		varStr = prefix + p.getName(true)
	} else {
		varStr = p.getName(false)
	}

	pointer := ""
	if def := p.getTypeDefinition(); def[0] == '*' { //is var definition has a pointer?
		pointer = "*"
	}
	convertStr := varStr

	if !p.isStringType() {
		convertStr = fmt.Sprintf("models.%sToString(%s%s)", p.TypeName(), pointer, varStr)
	}
	var paramAdder string

	if p.in != "path" {
		if p.in == "header" {
			paramAdder = fmt.Sprintf("request.AddHeader(\"%s\", %s)", p.id, convertStr)
		} else {
			paramAdder = fmt.Sprintf("request.AddParam(\"%s\", %s)", p.id, convertStr)
		}
	}
	if p.required { //required param
		if p.hasEmptyValue() {
			lines = append(lines, fmt.Sprintf("if len(%s) == 0 {\nerr = fmt.Errorf(\"%s is required\")\nreturn\n}", varStr, p.id))
		} else if !p.isPrimitive() { //struct type
			lines = append(lines, fmt.Sprintf("if %s == nil {\nerr = fmt.Errorf(\"%s is required\")\nreturn\n}", varStr, p.id))
		}
		lines = append(lines, paramAdder)
	} else { //optional param
		if p.hasEmptyValue() {
			lines = append(lines, fmt.Sprintf("if len(%s) > 0 {", varStr))
			lines = append(lines, paramAdder)
			lines = append(lines, fmt.Sprintf("}"))

		} else {
			lines = append(lines, fmt.Sprintf("if %s != nil {", varStr))
			lines = append(lines, paramAdder)
			lines = append(lines, fmt.Sprintf("}"))
		}
	}
	return strings.Join(lines, "\n")
}

func (p *Parameter) writeExtractParam(prefix string) string {
	out := []string{}

	var varStr string

	paramName := p.getName(false)

	out = append(out, fmt.Sprintf("\n// read '%s'", p.id))

	inStruct := len(prefix) > 0
	if inStruct {
		varStr = prefix + p.getName(true)
	} else {
		//write parameter definition
		out = append(out, fmt.Sprintf("var %s", p.getDefinition(false)))
		varStr = paramName
	}
	//get parameter string from request
	var convertFn, tempStr string

	if p.isStringType() {
		out = append(out, fmt.Sprintf("%s = ctx.Param(\"%s\")", varStr, p.id))
		tempStr = varStr
	} else {
		tempStr = paramName + "Str"
		convertFn = fmt.Sprintf("models.%sFromString(%s)", p.TypeName(), tempStr)
		out = append(out, fmt.Sprintf("%s := ctx.Param(\"%s\")", tempStr, p.id))
	}

	//check for non empty string and convert if neccessary
	if p.required {
		errBlock := fmt.Sprintf("\nresponse.SetBody(400, models.CreateProblemDetails(400, \"%s is required\"))\nreturn\n", p.id)
		out = append(out, fmt.Sprintf("if len(%s) == 0 {%s}\n", tempStr, errBlock))
		//convert from string
		if !p.isStringType() {
			errBlock = fmt.Sprintf("\nresponse.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf(\"parse %s failed: %%+v\", err)))\nreturn\n", p.id)
			out = append(out, fmt.Sprintf("if %s, err = %s; err != nil {%s}\n", varStr, convertFn, errBlock))
		}
	} else {
		//convert from string
		if !p.isStringType() {
			errBlock := fmt.Sprintf("\nresponse.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf(\"parse %s failed: %%+v\", err)))\nreturn\n", p.id)
			// optional primitive param needs a tmp var to keep converted value
			parsingBlock := ""
			convertVar := varStr
			if p.isPrimitive() {
				convertVar = paramName + "Tmp"
				parsingBlock = fmt.Sprintf("var %s %s\n", convertVar, p.m.goType)
			}

			parsingBlock = fmt.Sprintf("%sif %s, err = %s; err != nil {%s}\n", parsingBlock, convertVar, convertFn, errBlock)

			if p.isPrimitive() {
				parsingBlock = fmt.Sprintf("%s%s=&%s\n", parsingBlock, varStr, convertVar)
			}

			out = append(out, fmt.Sprintf("if len(%s) > 0 {\n%s}\n", tempStr, parsingBlock))
		}
	}

	return strings.Join(out, "\n")
}

func (p *Parameter) stringConvertFn(prefix string) string {
	var pStr string
	if len(prefix) > 0 {
		pStr = prefix + p.getName(true)
	} else {
		pStr = p.getName(false)
	}
	return pStr
}

type Operation struct {
	path        string
	pathTmpl    string
	pathParams  []string
	method      string
	id          string
	summary     string
	desc        string
	parameters  map[string]Parameter
	requestBody *RequestBody

	//for responses
	errorModel        *DataModel
	successModel      *DataModel
	errorCodes        []string //with Defined Error
	problemCodes      []string //with ProblemDetails
	successCode       string   //success code with a response (only one)
	emptySuccessCodes []string //success code with empty response
}

var models map[string]DataModel
var sbiRoot string = "sbi"
var sbiPackage string = "sbi"

func main() {
	specs := []string{
		"TS29503_Nudm_SDM.yaml",
		"TS29503_Nudm_UEAU.yaml",
		"TS29503_Nudm_UECM.yaml",
		"TS29507_Npcf_AMPolicyControl.yaml",
		"TS29512_Npcf_SMPolicyControl.yaml",
		"TS29525_Npcf_UEPolicyControl.yaml",
		"TS29502_Nsmf_PDUSession.yaml",
		//"TS29571_CommonData.yaml",
		"TS29518_Namf_Communication.yaml",
		"TS29509_Nausf_UEAuthentication.yaml",
	}
	for _, specFile := range specs {
		readSpec(specFile)
	}
}

func makeTitle(t string) string {
	return makeModelName(t)
}

func readSpec(specFile string) {
	config := &datamodel.DocumentConfiguration{
		AllowFileReferences:   true,
		AllowRemoteReferences: true,
		BasePath:              "specs",
	}

	specApis, _ := os.ReadFile("specs/" + specFile)
	// create a new document from specification bytes
	document, err := libopenapi.NewDocumentWithConfiguration(specApis, config)

	// if anything went wrong, an error is thrown
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	// because we know this is a v3 spec, we can build a ready to go model from it.
	v3Model, errors := document.BuildV3Model()

	// if anything went wrong when building the v3 model, a slice of errors will be returned
	if len(errors) > 0 {
		for i := range errors {
			log.Errorf("error: %+v", errors[i])
		}
		panic(fmt.Sprintf("cannot create v3 model from document: %d errors reported",
			len(errors)))
	}
	title := v3Model.Model.Info.Title
	title = makeTitle(title)
	log.Infof("Title: %s, Desc: %s", v3Model.Model.Info.Title, v3Model.Model.Info.Description)
	// get a count of the number of paths and schemas.
	paths := v3Model.Model.Paths.PathItems
	schemas := v3Model.Model.Components.Schemas

	log.Infof("There are %d paths and %d schemas in the document", paths.Len(), schemas.Len())
	createModels(schemas)
	operations := readPaths(paths)
	for _, m := range models {
		writeModel(&m)
	}
	var rootPath string
	if servers := v3Model.Model.Servers; len(servers) > 0 {
		rootPath = servers[0].URL
	}
	writeApis(title, rootPath, operations)
}

func writeEnum(id string, enum *Enum) {
	log.Infof("Define constant values for %s[%s]\n", id, enum.enumType)
	id = makeModelName(id)
	prefix := sbiRoot + "/models/"
	f, _ := os.Create(prefix + id + ".go")
	defer f.Close()

	writeFileHeader(f)

	fmt.Fprintf(f, "package models\n")

	fmt.Fprintf(f, "type %s %s\n", id, enum.enumType)
	fmt.Fprintf(f, "// Define constant values for %s\n", id)
	fmt.Fprintf(f, "const (\n")
	for _, v := range enum.values {
		v = strings.Replace(v, " ", "", -1)
		def := fmt.Sprintf("%s_%s", id, v)
		def = strings.Replace(def, "-", "_", -1)
		def = strings.ToUpper(def)
		if enum.enumType == "string" {
			fmt.Fprintf(f, "\t %s %s = \"%s\"\n", def, id, v)
		} else {
			fmt.Fprintf(f, "\t %s %s = %s\n", def, id, v)
		}
	}

	fmt.Fprintf(f, ") \n")
}
func writeModel(m *DataModel) {
	if m.enum != nil {
		writeEnum(m.id, m.enum)
		return
	}

	if isPrimitive(m.goType) {
		return
	}

	if len(m.properties) == 0 {
		return
	}

	structName := makeModelName(m.id)

	log.Infof("Write model %s", structName)

	prefix := sbiRoot + "/models/"

	f, _ := os.Create(prefix + structName + ".go")
	defer f.Close()
	writeFileHeader(f)
	fmt.Fprintf(f, "package models\n")

	fmt.Fprintf(f, "type %s struct {\n", structName)

	for _, p := range m.properties {
		if p.m == nil {
			log.Warnf("model %s has untype attribute %s", m.id, p.id)
			continue
		}

		goType := p.m.goType
		if p.m.enum != nil {
			goType = p.m.id
		}

		if !isPrimitive(goType) {
			goType = makeModelName(goType)
		}

		if p.m.isArray {
			goType = "[]" + goType
		} else if p.m.isMap {
			goType = "map[string]" + goType
		}

		attr := makeModelName(p.id)

		if p.required || p.hasEmptyValue() {
			fmt.Fprintf(f, "\t %s\t%s\t%s\n", attr, goType, p.writeTag())
		} else {
			fmt.Fprintf(f, "\t %s\t*%s\t%s\n", attr, goType, p.writeTag())
		}
	}
	fmt.Fprintf(f, "}\n")
}

func getRootPath(p string) string {
	parts := strings.Split(p, "/")
	if len(parts) > 1 {
		return strings.Join(parts[1:], "/")
	}
	return ""
}

func buildRoute(orgRoute string, params []string) (newRoute, routeTemplate string, foundParams []string, err error) {
	parts := strings.Split(orgRoute, "/")
	pathParams := []string{} //params in route
	newParts := []string{}   //for forming the route
	tmplParts := []string{}  //for forming route template

	for _, part := range parts {
		if param, ok := getPathParam(part); ok { //found a path param
			pathParams = append(pathParams, param)
			tmplParts = append(newParts, "%s")
			newParts = append(newParts, ":"+param)
		} else {
			newParts = append(newParts, part)
			tmplParts = append(tmplParts, part)
		}
	}
	if len(params) != len(pathParams) {
		err = fmt.Errorf("Num params not matched")
		return
	}

	for _, param := range params {
		if inList(param, pathParams) {
			foundParams = append(foundParams, param)
		} else {
			err = fmt.Errorf("param %s not in path", param)
			return
		}
	}

	newRoute = strings.Join(newParts, "/")
	routeTemplate = strings.Join(tmplParts, "/")
	return
}

func getPathParam(str string) (outParam string, ok bool) {
	if l := len(str); l >= 3 {
		if str[0] == '{' && str[l-1] == '}' {
			ok = true
			outParam = str[1 : l-1]
			return
		}
	}
	ok = false
	return
}

func writeApis(title string, rootPath string, operations map[string]Operation) {
	titleDir := sbiRoot + "/apis/" + title
	os.RemoveAll(titleDir)
	if err := os.Mkdir(titleDir, 0755); err != nil {
		log.Errorf("Fail to create Api directory for %s: %+v", title, err)
		return
	}
	fc, _ := os.Create(titleDir + "/consumer.go")
	defer fc.Close()
	writeFileHeader(fc)
	fmt.Fprintf(fc, "package %s\n", title)
	fmt.Fprintf(fc, "import (\n \"%s/models\"\n\"%s\"\n)\n", sbiPackage, sbiPackage)
	fmt.Fprintf(fc, "const (\n PATH_ROOT string = \"%s\"\n)\n", getRootPath(rootPath))

	for _, op := range operations {
		writeConsumerApi(fc, &op)
	}

	fp, _ := os.Create(titleDir + "/producer.go")
	defer fp.Close()
	writeFileHeader(fp)
	fmt.Fprintf(fp, "package %s\n", title)

	prodMethods := [][]byte{}
	prodMethodInfs := []string{}
	prodMethodImpls := []string{}
	useFmt := false
	for _, op := range operations {
		buf := new(bytes.Buffer)
		inf, impl, f := writeProducerApi(buf, &op, title)
		prodMethods = append(prodMethods, buf.Bytes())
		prodMethodInfs = append(prodMethodInfs, inf)
		prodMethodImpls = append(prodMethodImpls, impl)
		if f {
			useFmt = true
		}
	}
	if !useFmt {
		fmt.Fprintf(fp, "import (\n \"%s/models\"\n\"%s\"\n)\n", sbiPackage, sbiPackage)
	} else {
		fmt.Fprintf(fp, "import (\n\"fmt\"\n\"%s/models\"\n\"%s\"\n)\n", sbiPackage, sbiPackage)
	}
	for _, method := range prodMethods {
		fp.Write(method)
		fmt.Fprintf(fp, "\n\n")
	}

	fmt.Fprintf(fp, "type Producer interface {\n%s\n}\n", strings.Join(prodMethodInfs, "\n"))
	fi, _ := os.Create(titleDir + "/impl.go")
	defer fi.Close()
	fmt.Fprintf(fi, "package yourpkg\n\nimport \"%s/models\"\n\n/*\n%s\n*/", sbiPackage, strings.Join(prodMethodImpls, "\n\n"))
}

func writeFileHeader(f *os.File) {
	timeNow := time.Now().Format(time.UnixDate)
	fmt.Fprintf(f, "/*\nThis file is generated with a SBI APIs generator tool developed by ETRI\nGenerated at %v by TungTQ<tqtung@etri.re.kr>\nDo not modify\n*/\n\n", timeNow)
}

func writeProducerApi(f io.Writer, op *Operation, pkg string) (methodInf, methodImpl string, useFmt bool) {

	//write function definition
	fmt.Fprintf(f, "func On%s(ctx sbi.RequestContext, handler any) (response sbi.Response) {\n", op.id)
	fmt.Fprintf(f, "prod := handler.(Producer)\n")

	defineErr := false //should err be define?
	if op.requestBody != nil {
		defineErr = true
	} else {
		for _, p := range op.parameters {
			if !p.isStringType() {
				defineErr = true
			}
		}
	}

	if defineErr {
		fmt.Fprintf(f, "var err error\n")
		useFmt = true //use fmt package to write error
	}

	//write parameter extracting
	inputArgs := []string{}
	inputArgDefs := []string{}     //input argument definitions
	inputArgLongDefs := []string{} //input argument definitions with variables
	paramPrefix := ""

	if len(op.parameters) >= 2 {
		//write data structure to hold parameters
		paramStruct := op.id + "Params"
		fmt.Fprintf(f, "var params %s\n", paramStruct)
		inputArgs = append(inputArgs, "&params")
		inputArgDefs = append(inputArgDefs, "*"+paramStruct)
		inputArgLongDefs = append(inputArgLongDefs, fmt.Sprintf("params *%s.%s", pkg, paramStruct))
		paramPrefix = "params."
	} else {
		for _, p := range op.parameters {
			inputArgs = append(inputArgs, p.getName(false))
			inputArgDefs = append(inputArgDefs, p.getTypeDefinition())
			inputArgLongDefs = append(inputArgLongDefs, p.getDefinition(false))
		}
	}

	for _, p := range op.parameters {
		fmt.Fprintf(f, "%s\n", p.writeExtractParam(paramPrefix))
	}

	//write request decoding
	if op.requestBody != nil {
		fmt.Fprintf(f, "\n// decode request body\n")

		body := op.requestBody.content.goType
		if op.requestBody.required {
			fmt.Fprintf(f, "body := new(models.%s)\n", body)
			fmt.Fprintf(f, "if err = ctx.DecodeRequest(body); err != nil {\n")
			fmt.Fprintf(f, "response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf(\"Fail to decode request body: %%+v\", err)))\nreturn\n")
			fmt.Fprintf(f, "}\n")
		} else {
			fmt.Fprintf(f, "var body*models.%s\n", body)
			fmt.Fprintf(f, "if ctx.HaveRequestBody() {\n")

			fmt.Fprintf(f, "body = new(models.%s)\n", body)
			fmt.Fprintf(f, "if err = ctx.DecodeRequest(body); err !=nil {\n")
			fmt.Fprintf(f, "response.SetBody(400, models.NewSimpleProblem(400,fmt.Sprintf(\"Fail to decode request body: %%+v\", err)))\nreturn\n")
			fmt.Fprintf(f, "}\n")

			fmt.Fprintf(f, "}\n")
		}

		inputArgs = append(inputArgs, "body")
		inputArgDefs = append(inputArgDefs, "*models."+body)
		inputArgLongDefs = append(inputArgLongDefs, "body *models."+body)
	}
	//write calling application handler
	outputs := []string{}
	outputDefs := []string{}     //output definitions
	outputLongDefs := []string{} //output definitions with variables
	checkOutputs := []string{}

	if op.successModel != nil {
		outputs = append(outputs, "rsp")
		outputDefs = append(outputDefs, "*models."+op.successModel.goType)
		outputLongDefs = append(outputLongDefs, "rsp *models."+op.successModel.goType)
		tmp := fmt.Sprintf("\n// check for success response\n if rsp != nil {\nresponse.SetBody(%s, rsp)\nreturn\n}\n", op.successCode)
		checkOutputs = append(checkOutputs, tmp)
	}

	if len(op.errorCodes) > 0 && op.errorModel != nil {
		outputs = append(outputs, "ersp")
		outputDefs = append(outputDefs, "*"+op.errorModel.goType)
		outputLongDefs = append(outputLongDefs, "ersp *"+op.errorModel.goType)

		var errCode string
		if len(op.errorCodes) == 1 {
			errCode = op.errorCodes[0]
		} else {
			errCode = fmt.Sprintf("models.StatusFrom%s(ersp)", op.errorModel.goType)
		}
		tmp := fmt.Sprintf("\n// check for defined error\n if ersp != nil {\nresponse.SetBody(%s, ersp)\nreturn\n}\n", errCode)
		checkOutputs = append(checkOutputs, tmp)
	}
	if len(op.problemCodes) > 0 {
		outputs = append(outputs, "prob")
		outputDefs = append(outputDefs, "*models.ProblemDetails")
		outputLongDefs = append(outputLongDefs, "prob *models.ProblemDetails")
		tmp := fmt.Sprintf("\n // check for problem\n if prob != nil {\nresponse.SetBody(models.ProblemDetailsCode(prob), prob)\nreturn\n}\n")
		checkOutputs = append(checkOutputs, tmp)
	}

	fmt.Fprintf(f, "\n// call application handler\n")
	fmt.Fprintf(f, "%s := prod.Handle%s(%s)\n", strings.Join(outputs, ","), op.id, strings.Join(inputArgs, ","))

	methodImpl = fmt.Sprintf("func Handle%s(%s)(%s){\n\treturn\n}", op.id, strings.Join(inputArgLongDefs, ","), strings.Join(outputLongDefs, ","))
	methodInf = fmt.Sprintf("Handle%s(%s)(%s)\n", op.id, strings.Join(inputArgDefs, ","), strings.Join(outputDefs, ","))

	fmt.Fprintf(f, strings.Join(checkOutputs, ""))

	fmt.Fprintf(f, "return\n}\n")
	return
}

func writeConsumerApi(f *os.File, op *Operation) {
	fmt.Fprintf(f, "//Summary: %s\n", op.summary)
	fmt.Fprintf(f, "//Description: %s\n", op.desc)
	fmt.Fprintf(f, "//Path: %s\n", op.path)
	//fmt.Fprintf(f, "//Path Template: %s\n", op.pathTmpl)
	fmt.Fprintf(f, "//Path Params: %s\n", strings.Join(op.pathParams, ", "))

	inputArgs := []string{"cli sbi.ConsumerClient"}
	paramPrefix := ""

	if len(op.parameters) >= 2 {
		//write data structure to hold parameters
		paramStruct := fmt.Sprintf("%sParams", op.id)
		fmt.Fprintf(f, "type %s struct {\n", paramStruct)
		for _, p := range op.parameters {
			fmt.Fprintf(f, "%s\n", p.getDefinition(true))
		}
		fmt.Fprintf(f, "}\n")
		inputArgs = append(inputArgs, fmt.Sprintf("params %s", paramStruct))
		paramPrefix = "params."
	} else {
		for _, p := range op.parameters {
			inputArgs = append(inputArgs, p.getDefinition(false))
		}
	}

	if op.requestBody != nil {
		body := op.requestBody.content.id
		inputArgs = append(inputArgs, fmt.Sprintf("body *models.%s", body))
	}

	outputArgs := []string{}
	if op.successModel != nil {
		outputArgs = append(outputArgs, fmt.Sprintf("rsp *models.%s", op.successModel.goType))
	}

	if len(op.errorCodes) > 0 && op.errorModel != nil {
		outputArgs = append(outputArgs, fmt.Sprintf("ersp *models.%s", op.errorModel.goType))
	}

	outputArgs = append(outputArgs, "err error")
	//write function definition
	fmt.Fprintf(f, "func %s(%s) (%s) {\n", op.id, strings.Join(inputArgs, ","), strings.Join(outputArgs, ","))

	//write send request
	fmt.Fprintf(f, "\nrequest := sbi.DefaultRequest()\n")

	//write adding params to request
	paramChecks := []string{}
	for _, p := range op.parameters {
		if check := p.writeParamCheck(paramPrefix); len(check) > 0 {
			paramChecks = append(paramChecks, check)
		}
	}

	fmt.Fprintf(f, strings.Join(paramChecks, "\n"))

	//write check body required
	if op.requestBody != nil {
		if op.requestBody.required {
			fmt.Fprintf(f, "if body == nil {\n err = fmt.Errorf(\"body is required\")\nreturn\n}\n")
		}
		fmt.Fprintf(f, "request.Body = body\n")
	}

	//write path
	if len(op.pathParams) > 0 {
		pathParams := []string{}
		for _, pName := range op.pathParams {
			p, _ := op.parameters[pName]
			pathParams = append(pathParams, p.stringConvertFn(paramPrefix))
		}
		fmt.Fprintf(f, "\nrequest.Path= fmt.Sprintf(\"%%s%s\",PATH_ROOT, %s)\n", op.pathTmpl, strings.Join(pathParams, ", "))
	} else {
		fmt.Fprintf(f, "\nrequest.Path= fmt.Sprintf(\"%%s%s\",PATH_ROOT)\n", op.pathTmpl)
	}

	fmt.Fprintf(f, "var response sbi.Response\n")
	fmt.Fprintf(f, "if response, err = cli.Send(request); err !=nil {\n return\n}\n\n")

	//write check for response
	fmt.Fprintf(f, "switch response.StatusCode {\n")
	if op.successModel != nil {
		fmt.Fprintf(f, "case %s:\n", op.successCode)
		fmt.Fprintf(f, "rsp = new(%s)\n response.Body=rsp\nerr = cli.DecodeResponse(response)\n", op.successModel.goType)
	}
	if len(op.emptySuccessCodes) > 0 {
		fmt.Fprintf(f, "case %s:\nreturn\n", strings.Join(op.emptySuccessCodes, ","))
	}

	if len(op.errorCodes) > 0 && op.errorModel != nil {
		fmt.Fprintf(f, "case %s:\n", strings.Join(op.errorCodes, ","))
		fmt.Fprintf(f, "ersp = new(%s)\n response.Body=ersp\nerr = cli.DecodeResponse(response)\n", op.errorModel.goType)
	}

	if len(op.problemCodes) > 0 {
		fmt.Fprintf(f, "case %s:\n", strings.Join(op.problemCodes, ","))
		fmt.Fprintf(f, "prob := new(ProblemDetails)\n response.Body=prob\nif err = cli.DecodeResponse(response); err == nil {\nerr=sbi.ErrorFromProblemDetails(prob)\n}\n")
	}

	fmt.Fprintf(f, "default:\nerr=fmt.Errorf(\"%%d, %%s\",response.StatusCode, response.Status)\n}\n")
	fmt.Fprintf(f, "return\n}\n")
}

func readPaths(paths *orderedmap.Map[string, *v3.PathItem]) map[string]Operation {
	operations := make(map[string]Operation)
	for pathPair := paths.First(); pathPair != nil; pathPair = pathPair.Next() {
		// get the name of the schema
		pathName := pathPair.Key()

		// get the schema object from the map
		pathItem := pathPair.Value()
		op := readPathItem(pathName, pathItem)
		if op != nil {
			operations[op.id] = *op
		}
	}
	return operations
}

func createOpId(path string) string {
	parts := strings.FieldsFunc(path, func(c rune) bool {
		return c == '/' || c == '-' || c == '_' || c == ' '
	})
	idParts := []string{}
	for _, p := range parts {
		if len(p) > 0 && p[0] != '{' {
			idParts = append(idParts, strings.Title(p))
		}
	}
	return strings.Join(idParts, "")
}

func getRepresentativeContentModel(content *orderedmap.Map[string, *v3.MediaType]) *DataModel {
	if content == nil {
		return nil
	}
	var selectedModel *DataModel = nil
	for pair := content.First(); pair != nil; pair = pair.Next() {
		if m := analyzeSchema("", pair.Value().Schema); m != nil {
			if len(m.goType) == 0 {
				//return the first model created from inline schema
				return m
			} else if !isPrimitive(m.goType) {
				if selectedModel == nil {
					selectedModel = m
				} else {
					if m.goType != "ProblemDetails" {
						selectedModel = m
					}
				}
			}
		}
	}
	return selectedModel
}
func addNewModel(id string, m *DataModel) *DataModel {
	newM := new(DataModel)
	*newM = *m
	newM.id = id
	newM.goType = id
	models[id] = *newM
	return newM
}

func readPathItem(path string, item *v3.PathItem) *Operation {
	var op *v3.Operation
	var opStr string
	if item.Get != nil {
		op = item.Get
		opStr = "GET"
	} else if item.Put != nil {
		opStr = "PUT"
		op = item.Put
	} else if item.Post != nil {
		opStr = "POST"
		op = item.Post
	} else if item.Delete != nil {
		opStr = "DELETE"
		op = item.Delete
	} else if item.Patch != nil {
		opStr = "PATCH"
		op = item.Patch
	} else {
		log.Errorf("Unsupported operation on path %s", path)
		return nil
	}
	opModel := &Operation{
		method:     opStr,
		summary:    op.Summary,
		desc:       op.Description,
		parameters: make(map[string]Parameter),
	}
	//create operation id
	if len(op.OperationId) == 0 {
		opModel.id = createOpId(path + "/" + strings.Title(strings.ToLower(opStr)))
	} else {
		opModel.id = createOpId(op.OperationId)
	}

	//parse parameters
	parameters := make(map[string]*v3.Parameter)
	for _, p := range item.Parameters {
		parameters[p.Name] = p
	}
	for _, p := range op.Parameters {
		parameters[p.Name] = p
	}

	pathParams := []string{}
	for id, p := range parameters {
		var m *DataModel
		if m = analyzeSchema("", p.Schema); m == nil {
			//try to get it from content
			if p.Content != nil {
				for pair := p.Content.First(); pair != nil; pair = pair.Next() {
					if m = analyzeSchema("", pair.Value().Schema); m != nil {
						break
					}
				}
			}
		}

		if len(m.goType) == 0 {
			log.Errorf("Parameter %s of operation %s has inline-object type which is not supported", id, opModel.id)
			return nil
		}

		tmpP := Parameter{
			id:   id,
			m:    *m,
			in:   p.In,
			desc: p.Description,
		}
		if p.Required != nil {
			tmpP.required = *p.Required
		}
		opModel.parameters[tmpP.id] = tmpP
		if tmpP.in == "path" {
			pathParams = append(pathParams, tmpP.id)
		}
	}
	var err error
	if opModel.path, opModel.pathTmpl, opModel.pathParams, err = buildRoute(path, pathParams); err != nil {
		log.Errorf("Fail to build route template for the operation %s: %+v", opModel.id, err)
		return nil
	}

	//get request body
	if body := op.RequestBody; body != nil {
		if selectedModel := getRepresentativeContentModel(body.Content); selectedModel != nil {
			if len(selectedModel.goType) == 0 {
				//add model to the global repo
				modelId := fmt.Sprintf("%sRequest", opModel.id)
				selectedModel = addNewModel(modelId, selectedModel)
			}
			opModel.requestBody = &RequestBody{
				desc:    body.Description,
				content: *selectedModel,
			}
			if body.Required != nil {
				opModel.requestBody.required = *body.Required
			}
		}
	}

	if body := opModel.requestBody; body != nil {
		log.Infof("OP %s has body '%s[%v]'", opModel.id, body.content.id, body.required)
	} else {
		log.Infof("Op: '%s' has no request body", opModel.id)
	}

	//get responses
	for pair := op.Responses.Codes.First(); pair != nil; pair = pair.Next() {
		code := pair.Key()
		response := pair.Value()
		if code[0] == '1' {
			log.Warnf("respone with Information Http Code not processed")
			continue
		} else if code[0] == '3' {
			log.Warnf("respone with redirection Http Code not processed")
			continue
		} else if code[0] == '2' { //success response
			if selectedModel := getRepresentativeContentModel(response.Content); selectedModel != nil {
				if len(selectedModel.goType) == 0 {
					//assign model Id and goType then add to the model repo
					modelId := fmt.Sprintf("%sResponse", opModel.id)
					opModel.successModel = addNewModel(modelId, selectedModel)
				} else {
					opModel.successModel = selectedModel
				}
				opModel.successCode = code
			} else { //add success code with empty response
				opModel.emptySuccessCodes = append(opModel.emptySuccessCodes, code)
			}

		} else { //error responses
			if selectedModel := getRepresentativeContentModel(response.Content); selectedModel != nil {
				if len(selectedModel.goType) == 0 {
					//assign model Id and goType then add to the model repo
					modelId := fmt.Sprintf("%sErrorResponse", opModel.id)
					selectedModel = addNewModel(modelId, selectedModel)
				}

				if selectedModel.goType == "ProblemDetails" {
					opModel.problemCodes = append(opModel.problemCodes, code)
				} else {
					opModel.errorModel = selectedModel
					opModel.errorCodes = append(opModel.errorCodes, code)
				}
			}
		}
	}
	if opModel.successModel != nil {
		log.Infof("OP %s: success response:%s [%s]", opModel.id, opModel.successModel.goType, opModel.successCode)
	}
	log.Infof("OP %s: codes with problem details:%v", opModel.id, opModel.problemCodes)
	log.Infof("OP %s: codes with error response:%v", opModel.id, opModel.errorCodes)
	log.Infof("OP %s: codes with empty response:%v", opModel.id, opModel.emptySuccessCodes)

	return opModel
}

func showOperation(opStr string, op *v3.Operation) {
	for _, p := range op.Parameters {
		showParameter(p)
	}
	if op.RequestBody == nil {
		log.Infof("Operation has no request body")
	} else {
		showRequestBody(op.RequestBody)
	}
	if op.Responses == nil {
		log.Infof("Operation has no response")
	} else {
		showResponses(op.Responses)
	}

	if op.Callbacks != nil {
		for pair := op.Callbacks.First(); pair != nil; pair = pair.Next() {
			log.Infof("Callback for [%s]: %s", op.OperationId, pair.Key())
			showCallback(pair.Key(), pair.Value())
		}
	}

}

func showRequestBody(body *v3.RequestBody) {
	if body.Content != nil {
		num := 1
		for pair := body.Content.First(); pair != nil; pair = pair.Next() {
			if m := analyzeSchema("", pair.Value().Schema); m != nil {
				log.Infof("RequestBodyContent [%d] %s is %s[%s]", num, pair.Key(), m.goType, m.id)
				num++
			}
		}
	} else {
		log.Infof("Request: '%s' has no content", body.Description)
	}
}

func showResponses(responses *v3.Responses) {
	for pair := responses.Codes.First(); pair != nil; pair = pair.Next() {
		showResponse(pair.Key(), pair.Value())
	}
}

func showResponse(code string, response *v3.Response) {
	numContents := 0
	numHeaders := 0

	if response.Content != nil {
		for pair := response.Content.First(); pair != nil; pair = pair.Next() {
			if m := analyzeSchema("", pair.Value().Schema); m != nil {
				log.Infof("ResponseContent[%d] %s is %s[%s]", numContents+1, pair.Key(), m.goType, m.id)
				numContents++
			}
		}
	}

	if response.Headers != nil {
		numHeaders = response.Headers.Len()
	}
	log.Infof("Code=%s[content=%d][header=%d]", code, numContents, numHeaders)
}

func showParameter(p *v3.Parameter) {
	var schema *base.Schema
	if p.Schema == nil {
		log.Errorf("Parameter [%s] with no schema not supported", p.Name)
		return
	}

	if schema = p.Schema.Schema(); schema == nil {
		log.Errorf("Parameter [%s] with no schema not supported", p.Name)
		return
	}

	if p.Content != nil && p.Content.Len() > 0 {
		for pair := p.Content.First(); pair != nil; pair = pair.Next() {
			log.Infof("Parameter [%s] with content %s", p.Name, pair.Key())
		}
		log.Infof("Parameter [%s] with content not supported", p.Name)
		return
	}
	log.Infof("Parameter %s in %s", p.Name, p.In)
}

func showCallback(path string, callback *v3.Callback) {
	for pair := callback.Expression.First(); pair != nil; pair = pair.Next() {
		readPathItem(pair.Key(), pair.Value())
	}
}

func createModels(schemas *orderedmap.Map[string, *base.SchemaProxy]) {
	// print the number of paths and schemas in the document
	models = make(map[string]DataModel)

	for schemaPairs := schemas.First(); schemaPairs != nil; schemaPairs = schemaPairs.Next() {
		// get the name of the schema
		schemaName := schemaPairs.Key()

		// get the schema object from the map
		schemaValue := schemaPairs.Value()

		// build the schema
		schema := schemaValue.Schema()

		if schema == nil {
			log.Errorf("EMPTY SCHEMA '%s'\n", schemaName)
		} else {
			analyzeSchema(schemaName, schemaValue)
		}
	}
}
func getSchemaIdFromRef(ref string) string {
	tokens := strings.Split(ref, "/")
	if l := len(tokens); l > 0 {
		return tokens[l-1]
	}
	return ""
}

func analyzeAllOf(id string, allOf []*base.SchemaProxy) *DataModel {
	out := &DataModel{
		id:         id,
		goType:     id,
		properties: make(map[string]Property),
	}

	for _, s := range allOf {
		if m := analyzeSchema("", s); m != nil {
			for pId, p := range m.properties {
				out.properties[pId] = p
			}
		}
	}
	return out
}

func analyzeOneOf(id string, oneOf []*base.SchemaProxy) *DataModel {
	return analyzeAnyOf(id, oneOf)
}
func analyzeAnyOf(id string, anyOf []*base.SchemaProxy) *DataModel {

	isArray := false
	goTypes := make(map[string]DataModel)
	var enum *Enum
	var subModels []DataModel
	for _, s := range anyOf {
		if m := analyzeSchema("", s); m != nil {
			if m.isArray {
				isArray = true
			}
			if len(m.goType) > 0 {
				goTypes[m.goType] = *m
			}
			if m.enum != nil {
				enum = m.enum
			}
			subModels = append(subModels, *m)
		}
	}
	if len(goTypes) == 0 {
		return nil
	} else if len(goTypes) > 1 {
		if len(id) == 0 {
			return nil
		}

		log.Infof("ANYOF %s has more than one types", id)
		m := &DataModel{
			id:         id,
			goType:     id,
			properties: make(map[string]Property),
		}
		for tStr, t := range goTypes {
			if len(t.id) > 0 {
				tStr = makeModelName(t.id)
			}
			m.properties[tStr] = Property{
				id:       tStr,
				m:        &t,
				required: false,
			}
		}
		return m
	} else {
		if len(subModels) == 1 {
			m := subModels[0]
			if len(m.id) == 0 {
				m.id = id
			}
			return &m
		} else {
			m := &DataModel{
				isArray: isArray,
				id:      id,
				enum:    enum,
			}
			for t := range goTypes {
				m.goType = t
				break
			}
			return m
		}
	}
	return nil
}

func analyzeSchema(id string, schemaP *base.SchemaProxy) *DataModel {
	if schemaP == nil {
		return nil
	}
	if len(id) == 0 {
		if schemaP.IsReference() {
			id = getSchemaIdFromRef(schemaP.GetReference())
		}
	}
	if len(id) > 0 {
		if m, ok := models[id]; ok {
			return &m
		}
	}
	schema := schemaP.Schema()

	if schema == nil {
		return nil
	}

	if len(schema.Type) > 1 {
		log.Errorf("Schema with multiple types not supported")
		return nil
	}
	var m *DataModel
	if len(schema.Type) == 0 {
		if len(id) == 0 {
			if keyNode := schemaP.GetSchemaKeyNode(); keyNode != nil {
				id = keyNode.Value
			}
		}

		if len(schema.AllOf) > 0 {
			m = analyzeAllOf(id, schema.AllOf)
		} else if len(schema.AnyOf) > 0 {
			m = analyzeAnyOf(id, schema.AnyOf)
		} else if len(schema.OneOf) > 0 {
			m = analyzeOneOf(id, schema.OneOf)
		} else {
			//log.Warnf("UNTYPE schema %s", id)
			return nil
		}
	} else {
		m = &DataModel{
			id:         id,
			schema:     schema,
			properties: make(map[string]Property),
		}
		switch schema.Type[0] {
		case "object":
			if schema.AdditionalProperties != nil {
				m.goType = id
				m.isMap = true
				if schema.AdditionalProperties.IsA() {
					if refModel := analyzeSchema("", schema.AdditionalProperties.A); refModel != nil {
						m.goType = refModel.goType
					} else {
						m.goType = "UNKNOWN"
					}
				} else {
					m.goType = "bool"
				}
				//log.Infof("ADDPROP:map[string]%s", m.goType)
			} else {
				m.goType = id
				for pair := schema.Properties.First(); pair != nil; pair = pair.Next() {
					propName := pair.Key()
					propSchemaProxy := pair.Value()
					propSchema := propSchemaProxy.Schema()
					if propSchema == nil {
						log.Warnf("%s has empty attribute %s", id, propName)
					} else {
						p := Property{
							id:       propName,
							required: inList(propName, schema.Required),
						}

						if refModel := analyzeSchema("", propSchemaProxy); refModel != nil {
							p.m = refModel
						}
						m.properties[p.id] = p
					}
				}
				var extra *DataModel
				if l := len(schema.AllOf); l > 0 {
					extra = analyzeAllOf("", schema.AllOf)
				} else if l = len(schema.AnyOf); l > 0 {
					extra = analyzeAnyOf("", schema.AnyOf)
				} else if l = len(schema.OneOf); l > 0 {
					extra = analyzeOneOf("", schema.OneOf)
				}
				if extra != nil {
					for key, value := range extra.properties {
						m.properties[key] = value
					}
				}
			}
		case "string":
			if schema.Format == "binary" {
				m.goType = "byte"
				m.isArray = true
			} else {
				m.goType = "string"
			}
		case "integer":
			switch schema.Format {
			case "int32":
				m.goType = "int32"
			case "int64":
				m.goType = "int64"
			default:
				m.goType = "int"
			}
		case "number":
			switch schema.Format {
			case "double":
				m.goType = "float64"
			default:
				m.goType = "float64"
			}
		case "boolean":
			m.goType = "bool"

		case "array":
			m.isArray = true
			if schema.Items.IsA() {
				if itemModel := analyzeSchema("", schema.Items.A); itemModel != nil {
					m.goType = itemModel.goType
				} else {
					m.goType = "[]Unknown"
				}
			} else {
				m.goType = "bool"
			}

		default:
			log.Errorf("Not supported type: %s", schema.Type[0])
			return nil
		}
	}

	if len(schema.Enum) > 0 {
		m.enum = &Enum{
			enumType: m.goType,
		}
		for _, node := range schema.Enum {
			m.enum.values = append(m.enum.values, node.Value)
		}
		if len(id) == 0 {
			if keyNode := schemaP.GetSchemaKeyNode(); keyNode != nil {
				id = keyNode.Value
			}
		}
		m.id = id
	}

	if m != nil && len(m.id) > 0 {
		models[m.id] = *m
	}
	return m
}

var primitives []string = []string{"bool", "int", "int16", "int32", "int64", "float32", "float64", "string", "byte"}

func isPrimitive(t string) bool {
	return inList(t, primitives)
}

func indexInList(item string, list []string) int {
	for i, s := range list {
		if item == s {
			return i
		}
	}
	return -1
}

func inList(item string, list []string) bool {
	return indexInList(item, list) != -1
}

func makeModelName(s string) string {
	if len(s) == 0 {
		return ""
	}

	parts := strings.FieldsFunc(s, func(c rune) bool {
		return c == ' ' || c == '-' || c == '_'
	})
	out := ""
	for _, p := range parts {
		out = out + strings.Title(p)
	}

	if out[0] == '5' {
		out = "Five" + strings.Title(out[1:])
	}

	if out[0] == '3' {
		out = "Three" + strings.Title(out[1:])
	}

	return out
}
