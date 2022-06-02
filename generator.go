package docgen

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/swaggest/assertjson"
	"github.com/swaggest/openapi-go/openapi3"
)

func keepImportsOnSave() openapi3.Operation {
	ioutil.WriteFile("void", []byte{}, 0644)
	fmt.Errorf("")
	op := openapi3.Operation{}
	return op
}

type ApiDocument struct {
	Title   string
	Version string
}

func GenerateDoc(r *Router) error {
	// Read the Router.routes ang get  "name", "method", "query", "path", "body".
	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.1.0"}

	s := reflector.SpecEns() //  Création de la Spec (Object contenant les informations du service)

	for _, op := range r.routes {
		if &op != nil {

			oasOp := openapi3.Operation{}

			if strings.Contains(op.path, "{") {
				_path := op.path
				//var params []*openapi3.Parameter

				for strings.Contains(_path, "{") {
					_, aft, _ := strings.Cut(_path, "{")
					_path = aft
					pName, _, _ := strings.Cut(aft, "}")

					oasOp.Parameters = append(oasOp.Parameters,

						openapi3.ParameterOrRef{
							Parameter: &openapi3.Parameter{
								Name: pName,
								In:   openapi3.ParameterInPath,
							},
						},
					)
				}

			}
			// No need to "convert" the op.method on http.Method<type> as http.Method is nothing more than a string like "PATH", "GET", etc
			// BUT probably need to be uppercase
			reflector.SetRequest(&oasOp, op.input, strings.ToUpper(op.method)) // op = GET avec OrderItemRequestParams tant qu'input

			if op.resps != nil {
				for code, body := range op.resps {
					reflector.SetJSONResponse(&oasOp, body, code)
				}
			}

			s.AddOperation(op.method, op.path, oasOp)

		}
	}

	b, err := assertjson.MarshalIndentCompact(s, "", " ", 100) // Générer le resultat tant que doc JSON
	if err == nil {
		fmt.Print("NO ERROR")
	}
	ioutil.WriteFile("openapiROUTER.json", b, 0644) // écrire la doc résultat dans un file

	return nil
}
