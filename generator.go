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
			// No need to "convert" the op.method on http.Method<type> as http.Method is nothing more than a string like "PATH", "GET", etc
			// BUT probably need to be uppercase

			reflector.SetRequest(&oasOp, op.input, strings.ToUpper(op.method)) // op = GET avec OrderItemRequestParams tant qu'input

			//reflector.SetRequest(&oasOp, op.Input[string(meth)], http.MethodGet)      // op = GET avec OrderItemRequestParams tant qu'input
			//reflector.SetJSONResponse(&oasOp, op.Output[string(meth)], http.StatusOK) // Si OK (200), l'objet retour est du type OrderItemFilter
			// reflector.SetJSONResponse(&op, new([]OrderItemFilter), http.StatusConflict) // Si OK (409), l'objet retour est du type []OrderItemFilter

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
