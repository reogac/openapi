package main

import (
	"fmt"
	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/orderedmap"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
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

func (p *Property) isNullable() bool {
	if p.m.isArray {
		return true
	}

	if p.m.isMap {
		return true
	}

	if p.m.goType == "string" {
		return true
	}

	if p.m.goType == "[]byte" {
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

type Header struct {
	required bool
	m        DataModel
}

type RequestBody struct {
	required bool
	content  DataModel
	desc     string
}

type Response struct {
	headers map[string]Header
	content DataModel
}

type Parameter struct {
	id       string
	in       string
	required bool
	m        DataModel
}

type Operation struct {
	path        string
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

func main() {
	specs := []string{
		"TS29503_Nudm_SDM.yaml",
		"TS29503_Nudm_UEAU.yaml",
		"TS29503_Nudm_UECM.yaml",
		"TS29507_Npcf_AMPolicyControl.yaml",
		"TS29512_Npcf_SMPolicyControl.yaml",
		"TS29525_Npcf_UEPolicyControl.yaml",
		"TS29502_Nsmf_PDUSession.yaml",
		"TS29571_CommonData.yaml",
		"TS29518_Namf_Communication.yaml",
		"TS29509_Nausf_UEAuthentication.yaml",
	}
	for _, specFile := range specs {
		readSpec(specFile)
	}
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

	log.Infof("Title: %s, Desc: %s", v3Model.Model.Info.Title, v3Model.Model.Info.Description)
	// get a count of the number of paths and schemas.
	paths := v3Model.Model.Paths.PathItems
	schemas := v3Model.Model.Components.Schemas

	log.Infof("There are %d paths and %d schemas in the document", paths.Len(), schemas.Len())
	createModels(schemas)
	showPaths(paths)
	for _, m := range models {
		writeModel(&m)
	}
}

func writeEnum(id string, enum *Enum) {
	log.Infof("Define constant values for %s[%s]\n", id, enum.enumType)
	id = makeModelName(id)
	prefix := "models/"
	f, _ := os.Create(prefix + id + ".go")

	defer f.Close()

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

	prefix := "models/"
	f, _ := os.Create(prefix + structName + ".go")
	defer f.Close()

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

		if p.required || p.isNullable() {
			fmt.Fprintf(f, "\t %s\t%s\t%s\n", attr, goType, p.writeTag())
		} else {
			fmt.Fprintf(f, "\t %s\t*%s\t%s\n", attr, goType, p.writeTag())
		}
	}
	fmt.Fprintf(f, "}\n")
}

func showPaths(paths *orderedmap.Map[string, *v3.PathItem]) {
	for pathPair := paths.First(); pathPair != nil; pathPair = pathPair.Next() {
		// get the name of the schema
		pathName := pathPair.Key()

		// get the schema object from the map
		pathItem := pathPair.Value()
		showPathItem(pathName, pathItem)
	}
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

func showPathItem(path string, item *v3.PathItem) *Operation {
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
		path:       path,
		method:     opStr,
		summary:    op.Summary,
		desc:       op.Description,
		parameters: make(map[string]Parameter),
	}
	//create operation id
	if len(op.OperationId) == 0 {
		opModel.id = createOpId(opModel.path + "/" + strings.Title(strings.ToLower(opStr)))
	} else {
		opModel.id = createOpId(op.OperationId)
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
					opModel.successModel = selectedModel
					opModel.errorCodes = append(opModel.errorCodes, code)
				}
			}
		}
	}
	if opModel.successModel != nil {
		log.Infof("OP: success response:%s [%s]", opModel.successModel.goType, opModel.successCode)
	}
	log.Infof("OP: codes with problem details:%v", opModel.problemCodes)
	log.Infof("OP: codes with error response:%v", opModel.errorCodes)
	log.Infof("OP: codes with empty response:%v", opModel.emptySuccessCodes)
	/*
		for _, p := range item.Parameters {
			showParameter(p)
		}
	*/
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
		showPathItem(pair.Key(), pair.Value())
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
				m.goType = "[]byte"
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

var primitives []string = []string{"bool", "int", "int16", "int32", "int64", "float32", "float64", "string", "[]byte"}

func isPrimitive(t string) bool {
	return inList(t, primitives)
}

func inList(item string, list []string) bool {
	for _, s := range list {
		if item == s {
			return true
		}
	}
	return false
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
