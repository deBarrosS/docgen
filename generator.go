package docgen

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/swaggest/assertjson"
	"github.com/swaggest/openapi-go/openapi3"
)

func GenerateDoc(r *Router) error {
	// Read the Router.routes ang get  "name", "method", "query", "path", "body".
	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.1.0"}

	s := reflector.SpecEns() //  Création de la Spec (Object contenant les informations du service)

	// o(nRoutes*(nParams+nResps))
	for _, op := range r.routes {
		if &op != nil {
			oasOp := openapi3.Operation{}

			if &op.path != nil && strings.Contains(op.path, "{") { //o(len(op.path))
				op.setParameters(&oasOp)
			}

			reflector.SetRequest(&oasOp, op.input, strings.ToUpper(op.method))

			if op.resps != nil {
				for code, body := range op.resps {
					reflector.SetJSONResponse(&oasOp, body, code)
				}
			}

			if op.internal {
				oasOp.MapOfAnything = map[string]interface{}{
					"x-internal": true,
				}
			}
			oasOp.Deprecated = &op.deprecated

			s.AddOperation(op.method, op.path, oasOp)
		}
	}

	b, err := assertjson.MarshalIndentCompact(s, "", " ", 100) // Générer le resultat tant que doc JSON
	if err == nil {
		fmt.Print("NO ERROR")
	}
	ioutil.WriteFile("openapiHandler.json", b, 0644) // écrire la doc résultat dans un file

	return nil
}

// o(nParams)
func (r *Route) setParameters(oasOp *openapi3.Operation) {
	_path := r.path
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
			// o(1)
			oasOp.Parameters[n] = openapi3.ParameterOrRef{
				Parameter: &openapi3.Parameter{
					Name: pName,
					In:   openapi3.ParameterInPath,
				},
			}
		}
	}
}
