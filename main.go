package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/orderedmap"
)

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
	//specApis, _ := os.ReadFile("specs/TS29502_Nsmf_PDUSession.yaml")
	specApis, _ := os.ReadFile("specs/TS29571_CommonData.yaml")
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
	showSchemas(schemas)
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

func showSchemas(schemas *orderedmap.Map[string, *base.SchemaProxy]) {
	// print the number of paths and schemas in the document

	for schemaPairs := schemas.First(); schemaPairs != nil; schemaPairs = schemaPairs.Next() {
		// get the name of the schema
		schemaName := schemaPairs.Key()

		// get the schema object from the map
		schemaValue := schemaPairs.Value()

		// build the schema
		schema := schemaValue.Schema()

		// if the schema has properties, print the number of
		// properties
		if schema == nil {
			fmt.Printf("EMPTY SCHEMA '%s'\n", schemaName)
		} else {
			printSchema(schemaName, schema)

		}
	}
}

func getSchemaType(schemaName string, schema *base.Schema) string {
	if schema.Properties != nil {
		return schema.SchemaTypeRef
	}
	if len(schema.Type) == 0 {
		return "Unknown"
	}
	switch schema.Type[0] {
	case "array":
		if schema.Items.IsA() {
			itemSchema := schema.Items.A.Schema()
			json, _ := itemSchema.MarshalJSONInline()
			fmt.Printf("%s\n", string(json))
			if len(itemSchema.Type) > 0 {
				if len(itemSchema.SchemaTypeRef) > 0 {
					return "[]" + itemSchema.SchemaTypeRef
				} else {
					return "[]" + itemSchema.Type[0]
				}
			} else {
				return "[]" + itemSchema.SchemaTypeRef
			}
		} else {
			return "[]bool"
		}
	default:
		return schema.Type[0]
	}
}

func inList(item string, list []string) bool {
	for _, s := range list {
		if item == s {
			return true
		}
	}
	return false
}
func prettyJson(input []byte) []byte {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "    ")
	if err := enc.Encode(input); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
func printSchema(schemaName string, schema *base.Schema) {
	prefix := "models/"

	f, _ := os.Create(prefix + schemaName + ".json")
	defer f.Close()
	json, _ := schema.MarshalJSONInline()
	fmt.Fprintf(f, "%s", json)
	return

	if schema.Nullable != nil {
		fmt.Printf("SCHEMA NULLABLE'%s'\n", schemaName)
	}

	if numTypes := len(schema.Type); numTypes == 0 {
		fmt.Printf("SCHEMA NO-TYPE'%s'\n", schemaName)
		if len(schema.AnyOf) > 0 {
			fmt.Printf("SCHEMA ANYOF'%s'\n", schemaName)
			for _, anyOf := range schema.AnyOf {
				subSchema := anyOf.Schema()
				//fmt.Printf("SUBSCHEMA='%s'\n", subSchema.Type[0])
				if len(subSchema.Enum) > 0 {
					for _, v := range subSchema.Enum {
						fmt.Printf("ENUM'%s':%s\n", schemaName, v.Value)
					}
				}
			}
		}

		if len(schema.OneOf) > 0 {
			fmt.Printf("SCHEMA ONEOF'%s':%v\n", schemaName, schema.OneOf)
		}
		if len(schema.AllOf) > 0 {
			fmt.Printf("SCHEMA ALLOF'%s':%v\n", schemaName, schema.AllOf)
		}

	} else if numTypes == 1 {
		//fmt.Printf("SCHEMA '%s' has type:'%s'[%d]\n", schemaName, schema.Type[0], len(schema.Type))
		if schema.Properties != nil { //object type

			f, _ := os.Create(prefix + schemaName + ".go")
			defer f.Close()

			//fmt.Printf("Schema '%s' has %d properties\n", schemaName, schema.Properties.Len())
			//fmt.Printf("%v\n", schema)
			fmt.Fprintf(f, "type %s struct {\n", schemaName)
			for pair := schema.Properties.First(); pair != nil; pair = pair.Next() {
				probName := pair.Key()
				probSchema := pair.Value().Schema()
				if probSchema == nil {
					fmt.Printf("%s has empty attribute %s\n", schemaName, probName)
				} else {
					if inList(probName, schema.Required) {
						fmt.Fprintf(f, "\t%s\t%s\n", probName, getSchemaType(schemaName, probSchema))
					} else {
						fmt.Fprintf(f, "\t%s\t*%s\n", probName, getSchemaType(schemaName, probSchema))
					}
				}
			}
			fmt.Fprintf(f, "}\n")
		}
	} else {
		fmt.Printf("SCHEMA '%s' has multiple types[%d]\n", schemaName, numTypes)
	}

}
