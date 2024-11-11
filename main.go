package main

import (
	"fmt"
	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/orderedmap"
	"os"
	"strings"
)

type DataModel struct {
	id          string
	goType      string
	isPrimitive bool
	isArray     bool
	properties  map[string]Property
	schema      *base.Schema
}

type Property struct {
	id       string
	required bool
	goType   string
	isArray  bool
}

func (p *Property) isNullable() bool {
	if p.isArray {
		return true
	}

	if p.goType == "string" {
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

var models map[string]DataModel

func main() {
	readSpec()
}

func readSpec() {
	config := &datamodel.DocumentConfiguration{
		AllowFileReferences:   true,
		AllowRemoteReferences: true,
		BasePath:              "specs",
	}
	// load an OpenAPI 3 specification from bytes
	specApis, _ := os.ReadFile("specs/TS29502_Nsmf_PDUSession.yaml")
	//specApis, _ := os.ReadFile("specs/TS29571_CommonData.yaml")
	//specApis, _ := os.ReadFile("specs/TS29518_Namf_Communication.yaml")

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
			fmt.Printf("error: %e\n", errors[i])
		}
		panic(fmt.Sprintf("cannot create v3 model from document: %d errors reported",
			len(errors)))
	}

	fmt.Printf("Title: %s, Desc: %s\n", v3Model.Model.Info.Title, v3Model.Model.Info.Description)
	// get a count of the number of paths and schemas.
	paths := v3Model.Model.Paths.PathItems
	schemas := v3Model.Model.Components.Schemas

	fmt.Printf("There are %d paths and %d schemas in the document", paths.Len(), schemas.Len())
	//showPaths(paths)
	createModels(schemas)
	for _, m := range models {
		writeModel(&m)
	}

}

func writeModel(m *DataModel) {
	if m.isPrimitive {
		return
	}
	if len(m.properties) == 0 {
		return
	}

	prefix := "models/"
	f, _ := os.Create(prefix + m.id + ".go")
	defer f.Close()
	fmt.Fprintf(f, "type %s struct {\n", m.id)

	for _, p := range m.properties {
		goType := p.goType
		if len(goType) == 0 {
			fmt.Printf("model %s has untype attribute %s\n", m.id, p.id)
			continue
		}
		if p.isArray {
			goType = "[]" + goType
		}
		if p.required || p.isNullable() {
			fmt.Fprintf(f, "\t %s\t%s\t%s\n", strings.Title(p.id), goType, p.writeTag())
		} else {
			fmt.Fprintf(f, "\t %s\t*%s\t%s\n", strings.Title(p.id), goType, p.writeTag())
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
func showPathItem(path string, item *v3.PathItem) {
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
		fmt.Printf("Unsupported operation\n")
		return
	}
	fmt.Printf("%s[%s][%s]: %s\n", path, opStr, op.OperationId, op.Summary)

	for _, p := range op.Parameters {
		showParameter(p)
	}

	if op.Responses == nil {
		fmt.Printf("Operation has no response\n")
	} else {
		showResponses(op.Responses)
	}

	if op.Callbacks != nil {
		for pair := op.Callbacks.First(); pair != nil; pair = pair.Next() {
			fmt.Printf("Callback for [%s]: %s\n", op.OperationId, pair.Key())
			showCallback(pair.Key(), pair.Value())
		}
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
		numContents = response.Content.Len()
	}

	if response.Headers != nil {
		numHeaders = response.Headers.Len()
	}
	fmt.Printf("Code=%s[content=%d][header=%d]\n", code, numContents, numHeaders)
}

func showParameter(p *v3.Parameter) {
	var schema *base.Schema
	if p.Schema == nil {
		fmt.Printf("Parameter [%s] with no schema not supported\n", p.Name)
		return
	}

	if schema = p.Schema.Schema(); schema == nil {
		fmt.Printf("Parameter [%s] with no schema not supported\n", p.Name)
		return
	}

	if p.Content != nil && p.Content.Len() > 0 {
		for pair := p.Content.First(); pair != nil; pair = pair.Next() {
			fmt.Printf("Parameter [%s] with content %s\n", p.Name, pair.Key())
		}
		fmt.Printf("Parameter [%s] with content not supported\n", p.Name)
		return
	}
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
			fmt.Printf("EMPTY SCHEMA '%s'\n", schemaName)
		} else {
			analyzeSchema(schemaValue)
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
		if m := analyzeSchema(s); m != nil {
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
	goTypes := make(map[string]bool)
	for _, s := range anyOf {
		if m := analyzeSchema(s); m != nil {
			if m.isArray {
				isArray = true
			}
			if len(m.goType) > 0 {
				goTypes[m.goType] = true
			}
		}
	}
	if len(goTypes) == 0 {
		return nil
	} else if len(goTypes) > 1 {
		fmt.Printf("ANYOF %s has more than one types\n", id)
		return nil
	} else {
		m := &DataModel{
			isArray: isArray,
			id:      id,
		}
		for t, _ := range goTypes {
			m.goType = t
			break
		}
		return m
	}
	return nil
}
func getSchemaId(schemaP *base.SchemaProxy) string {
	if schemaP.IsReference() {
		return getSchemaIdFromRef(schemaP.GetReference())
	} else if keyNode := schemaP.GetSchemaKeyNode(); keyNode != nil {
		return keyNode.Value
	}

	return ""
}

func analyzeSchema(schemaP *base.SchemaProxy) *DataModel {
	id := getSchemaId(schemaP)
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
		fmt.Printf("Schema with multiple types not supported\n")
		return nil
	}
	var m *DataModel
	if len(schema.Type) == 0 {
		if len(schema.AllOf) > 0 {
			m = analyzeAllOf(id, schema.AllOf)
		} else if len(schema.AnyOf) > 0 {
			m = analyzeAnyOf(id, schema.AnyOf)
		} else if len(schema.OneOf) > 0 {
			m = analyzeOneOf(id, schema.OneOf)
		} else {
			fmt.Printf("UNTYPE %s\n", id)
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
			m.goType = id
			for pair := schema.Properties.First(); pair != nil; pair = pair.Next() {
				propName := pair.Key()
				propSchemaProxy := pair.Value()
				propSchema := propSchemaProxy.Schema()
				if propSchema == nil {
					fmt.Printf("%s has empty attribute %s\n", id, propName)
				} else {
					p := Property{
						id:       propName,
						required: inList(propName, schema.Required),
					}

					if refModel := analyzeSchema(propSchemaProxy); refModel != nil {
						p.goType = refModel.goType
						p.isArray = refModel.isArray

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
				extra = analyzeAnyOf("", schema.OneOf)
			}
			if extra != nil {
				for key, value := range extra.properties {
					m.properties[key] = value
				}
			}
		case "string":
			m.goType = "string"
			m.isPrimitive = true
		case "integer":
			m.isPrimitive = true
			switch schema.Format {
			case "int32":
				m.goType = "int32"
			case "int64":
				m.goType = "int64"
			default:
				m.goType = "int"
			}
		case "number":
			m.isPrimitive = true
			switch schema.Format {
			case "double":
				m.goType = "float64"
			default:
				m.goType = "float"
			}
		case "boolean":
			m.goType = "bool"
			m.isPrimitive = true

		case "array":
			m.isPrimitive = false
			m.isArray = true
			if schema.Items.IsA() {
				if itemModel := analyzeSchema(schema.Items.A); itemModel != nil {
					m.goType = itemModel.goType
				} else {
					m.goType = "[]Unknown"
				}
			} else {
				m.goType = "bool"
			}

		default:
			fmt.Printf("Not supported type: %s\n", schema.Type[0])
			return nil
		}
		fmt.Printf("%s: %s\n", id, m.goType)
	}
	if m != nil && len(m.id) > 0 {
		models[m.id] = *m
	}
	return m
}

func inList(item string, list []string) bool {
	for _, s := range list {
		if item == s {
			return true
		}
	}
	return false
}
