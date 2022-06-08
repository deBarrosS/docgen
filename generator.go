package docgen

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/swaggest/openapi-go/openapi3"
)

func GenerateDoc(r *Router, name string) error {
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

	schema, err := reflector.Spec.MarshalYAML()

	if err == nil {
		fmt.Print("NO ERROR")
	}
	ioutil.WriteFile("generated/openapi"+name+".yaml", schema, 0644) // écrire la doc résultat dans un file

	return nil
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
