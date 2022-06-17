package docgen

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/swaggest/openapi-go/openapi3"
)

func GenerateDoc(r *Router, name string) ([]byte, error) {
	// Read the Router.routes ang get  "name", "method", "query", "path", "body".
	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.1.0"}

	s := reflector.SpecEns() //  Création de la Spec (Object contenant les informations du service)

	// o(nRoutes*(nParams+nResps))
	for _, op := range r.Routes {
		if &op != nil {
			oasOp := openapi3.Operation{}

			if &op.Path != nil && strings.Contains(op.Path, "{") { //o(len(op.Path))
				op.setParameters(&oasOp)
			}

			reflector.SetRequest(&oasOp, op.Input, strings.ToUpper(op.Method))

			if op.Resps != nil {
				for code, body := range op.Resps {
					reflector.SetJSONResponse(&oasOp, body, code)
				}
			}

			if op.Internal {
				oasOp.MapOfAnything = map[string]interface{}{
					"x-internal": true,
				}
			}
			oasOp.Deprecated = &op.Deprecated

			s.AddOperation(op.Method, op.Path, oasOp)
		}
	}
	for _, comp := range s.Components.Schemas.MapOfSchemaOrRefValues {
		setInternal(&comp, s)
	}
	schema, err := reflector.Spec.MarshalYAML()

	if err == nil {
		fmt.Print("NO ERROR")
	}

	return schema, nil
}

func GenerateDocWrite(r *Router, name string) ([]byte, error) {
	// Read the Router.routes ang get  "name", "method", "query", "path", "body".
	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.1.0"}

	s := reflector.SpecEns() //  Création de la Spec (Object contenant les informations du service)
	internalTag := "[internal]"

	// o(nRoutes*(nParams+nResps))
	for _, op := range r.Routes {
		if &op != nil {
			oasOp := openapi3.Operation{}

			if &op.Path != nil && strings.Contains(op.Path, "{") { //o(len(op.Path))
				op.setParameters(&oasOp)
			}

			reflector.SetRequest(&oasOp, op.Input, strings.ToUpper(op.Method))

			if oasOp.RequestBody != nil {
				if oasOp.RequestBody.RequestBody.Content["application/json"].Schema.Schema != nil {
					oasOp.RequestBody.RequestBody.Content["application/json"].Schema.Schema.WithMapOfAnythingItem("x-internal", true)
				}
				if oasOp.RequestBody.RequestBody.Content["application/json"].Schema.SchemaReference != nil {
					ref := oasOp.RequestBody.RequestBody.Content["application/json"].Schema.SchemaReference.Ref
					ref = strings.TrimPrefix(ref, "#/components/schemas/")
					sch := s.Components.Schemas.MapOfSchemaOrRefValues[ref].Schema
					sch.WithMapOfAnythingItem("x-internal", true)
					print("debug")
				}

				//oasOp.RequestBody.RequestBody.MapOfAnything["x-internal"] = true
			}

			if op.Resps != nil {
				for code, body := range op.Resps {
					reflector.SetJSONResponse(&oasOp, body, code)
				}
			}
			oasOp.MapOfAnything = map[string]interface{}{
				"x-internal": true,
			}

			oasOp.Deprecated = &op.Deprecated
			oasOp.Description = &internalTag

			s.AddOperation(op.Method, op.Path, oasOp)
		}
	}

	//o(comp * props)
	for _, comp := range s.Components.Schemas.MapOfSchemaOrRefValues {
		setInternal(&comp, s)
	}

	schema, err := reflector.Spec.MarshalYAML()

	ioutil.WriteFile("openapi.yaml", schema, 0644)

	if err == nil {
		fmt.Print("NO ERROR")
	}

	return schema, nil
}

// Need oasSpec to find the components referenced by others
// o(props)
func setInternal(s *openapi3.SchemaOrRef, oasSpec *openapi3.Spec) {
	tagInternal := "[internal]"

	// o(1)
	if s.SchemaReference != nil {
		ref := s.SchemaReference.Ref
		if &ref != nil {
			refTrim := strings.TrimPrefix(ref, "#/components/schemas/")
			sch := oasSpec.Components.Schemas.MapOfSchemaOrRefValues[refTrim]
			if &sch != nil {
				setInternal(&sch, oasSpec)
			}
		}
	}
	// o(props)
	if s.Schema != nil {
		s.Schema.WithMapOfAnythingItem("x-internal", true)
		s.Schema.WithDescription(tagInternal)

		for _, prop := range s.Schema.Properties {
			if prop.Schema != nil {
				prop.Schema.WithMapOfAnythingItem("x-internal", true)
				prop.Schema.WithDescription(tagInternal)
			}
		}
	}
}

// o(nParams)
func (r *Route) setParameters(oasOp *openapi3.Operation) {
	_path := r.Path
	n := strings.Count(_path, "{") //o(len(_path))

	oasOp.Parameters = make([]openapi3.ParameterOrRef, n)

	//o(n + 2*strings.Cut)
	for n >= 0 {
		n--

		// Get param name between '{' et '}'
		_, aft, _ := strings.Cut(_path, "{")
		_path = aft
		pName, _, ok := strings.Cut(aft, "}")

		if ok {
			required := true
			strSchema := openapi3.SchemaTypeString

			// o(1)
			oasOp.Parameters[n] = openapi3.ParameterOrRef{
				Parameter: &openapi3.Parameter{
					Name:     pName,
					In:       openapi3.ParameterInPath,
					Required: &required,
					Schema: &openapi3.SchemaOrRef{
						Schema: &openapi3.Schema{
							Type: &strSchema,
						},
					},
				},
			}
		}
	}
}
