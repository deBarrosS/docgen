package old

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swaggest/assertjson"
	"github.com/swaggest/openapi-go/openapi3"
)

type Api struct {
	Name      string
	Version   string
	AllRoutes []*Operation
	NbRoutes  int
}

type Operation struct {
	Route  *mux.Route
	Input  map[string]interface{}
	Output map[string]interface{}
}

// Initialize the specification of the API. The number of routes is default to 30.
func (api *Api) InitApi(name, version string, nbRoutes int) *Api {
	api.Name = name
	api.Version = version

	if len(name) <= 0 {
		api.Name = "{Change API Name}"
	}

	if len(version) <= 0 {
		api.Version = "{Change API Version}"
	}

	if nbRoutes <= 0 {
		api.AllRoutes = make([]*Operation, 30)
	} else {
		api.AllRoutes = make([]*Operation, nbRoutes)
	}

	return api
}

// As we are returning pointers, no problem regarding the changes to the struct
func (api *Api) Add(newRoute *mux.Route) *Operation {
	if api == nil {
		api = new(Api)
		// AllRoutes is not an array but a slice therefore it needs to be initialized this way
		api.AllRoutes = make([]*Operation, 10)
	}

	op := new(Operation)
	op.Route = newRoute

	if api.AllRoutes == nil {
		return nil //, errors.New("Api not initialized")
	}

	if api.NbRoutes < cap(api.AllRoutes) {
		api.AllRoutes[api.NbRoutes] = op
		api.NbRoutes++
	}

	return op //, nil
}

func (op *Operation) BodyInput(meth string, body interface{}) *Operation {
	if op == nil {
		return nil //, errors.New("Operation not initialized")
	}
	if op.Input == nil {
		op.Input = make(map[string]interface{})
	}
	op.Input[meth] = body

	return op //, nil
}

func (op *Operation) BodyOutput(respCode string, body interface{}) (*Operation, error) {
	if op == nil {
		return nil, errors.New("Operation not initialized")
	}

	if op.Output == nil {
		op.Output = make(map[string]interface{})
	}

	// handle overwiting a response body?
	op.Output[respCode] = body

	return op, nil
}

func (api *Api) GenerateDoc() ([]byte, error) {
	if api == nil {
		return nil, errors.New("API not initialized")
	}
	// Iterate through the api specification and add the operations to the openapi doc
	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.0.3"}

	s := reflector.SpecEns() //  Création de la Spec (Object contenant les informations du service)
	s.Info.Title = api.Name  //  Set le Titre de la documentation/ du document openapi résultat
	s.Info.Version = api.Version

	for _, op := range api.AllRoutes {
		methods, err := op.Route.GetMethods()
		if err == nil {
			for _, meth := range methods {
				oasOp := openapi3.Operation{}

				input := op.Input[string(meth)]

				path, err := op.Route.GetPathTemplate()
				method := http.MethodGet
				if err != nil {
					continue
				}
				switch meth {
				case "PATCH":
					method = http.MethodPatch
					break
				case "POST":
					method = http.MethodPost
					break
				case "DELETE":
					method = http.MethodDelete
					break
				case "PUT":
					method = http.MethodPut
					break
				}
				reflector.SetRequest(&oasOp, input, method) // op = GET avec OrderItemRequestParams tant qu'input
				//reflector.SetRequest(&oasOp, op.Input[string(meth)], http.MethodGet)      // op = GET avec OrderItemRequestParams tant qu'input
				reflector.SetJSONResponse(&oasOp, op.Output[string(meth)], http.StatusOK) // Si OK (200), l'objet retour est du type OrderItemFilter
				// reflector.SetJSONResponse(&op, new([]OrderItemFilter), http.StatusConflict) // Si OK (409), l'objet retour est du type []OrderItemFilter
				s.AddOperation(method, path, oasOp)
			}
		}
	}

	b, err := assertjson.MarshalIndentCompact(s, "", " ", 100) // Générer le resultat tant que doc JSON
	if err == nil {
		fmt.Print("NO ERROR")
	}
	ioutil.WriteFile("openapi.json", b, 0644) // écrire la doc résultat dans un file

	return []byte("TODO"), nil
}
