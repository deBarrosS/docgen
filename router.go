package docgen

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	//1router     *mux.Router
	routes     []*Route
	nbRoutes   int
	default400 interface{}
	default500 interface{} // etc...

}

type Route struct {
	//1route *mux.Route             // Each route creation is responsible of creating a whole new mux.Route
	path    string
	method  string
	input   interface{}            // No need to map by method because of the line above ^^
	resps   map[string]interface{} // outputs per return code
	handler func()
}

func InitApi(name, version string) *Router {
	r := new(Router)
	r.routes = make([]*Route, 30)
	r.nbRoutes = 0
	// r.doc = new(Api)
	// r.doc.InitApi(name, version, 30)

	return r
}

func (r *Router) NewRoute() *Route {

	nR := new(Route)
	// by attaching the *mux.Route to the Route struct we'll be able to modify the real mux.Route

	//1nR.route = r.router.NewRoute()

	if r.nbRoutes < len(r.routes) {
		r.routes[r.nbRoutes] = nR
		r.nbRoutes++
	}
	// else grow slice

	return nR
}

func (r *Route) Path(path string) *Route {
	// call the attachement of the path to the mux.Route
	// As it is r is a pointer to a struct that contains a pointer to a mux.Route,
	// the changes would be reflected on the mux.Routes attached to the mux.Router

	//1r.route.Path(path)
	r.path = path

	// return our strcuture
	return r
}

func (r *Route) Methods(meth string) *Route {

	//1r.route.Methods(meth)
	r.method = meth
	return r
}

func (r *Route) Method(meth string, input interface{}) *Route {

	r.input = input
	r.method = meth
	//1r.route.Methods(meth)
	return r
}

// Possible to change the function signature to adapt to our needs without having to stop using mux
func (r *Route) Handler(handler http.Handler, input, output string) *Route {
	/*--------------- SI Methods(meth, input)
	Lors de l'attribution d'un handler à une route on disposera du format de l'input attendu. (r.input)
	We could then pass it ahead to the handler, the validator, etc... (If needed)

	Therefore, no need to pass it here as params

	!!!!!!! Need to call Methods before calling Handler though !!!!
	*/

	//1r.route.Handler(handler)

	return r
}

func (r *Route) ExpectedResponses(code string, body interface{}) *Route {

	r.resps[code] = body

	return r
}

// Router as param or as caller ?
// Pass "app" as parameter ?
func (r *Router) RegisterRoutes() (*mux.Router, error) {
	router := new(mux.Router)

	for _, route := range r.routes {
		router.NewRoute().Path(route.path).Methods(route.method) //Add handler
		// Shall we add some error handling here? If an error, then log (the x route couldn't be set and continue ?)
	}

	return router, nil
}

// func (r *Router) GenerateDoc() error {
// 	// Read the Router.routes ang get  "name", "method", "query", "path", "body".
// 	reflector := openapi3.Reflector{}
// 	reflector.Spec = &openapi3.Spec{Openapi: "3.1.0"}

// 	s := reflector.SpecEns()  //  Création de la Spec (Object contenant les informations du service)
// 	//s.Info.Title = r.doc.Name //  Set le Titre de la documentation/ du document openapi résultat
// 	//s.Info.Version = r.doc.Version

// 	for _, op := range r.routes {
// 		if op != nil {
// 			methods, err := "122","222"//1op.route.GetMethods()
// 			if err == nil {
// 				for _, meth := range methods {
// 					oasOp := openapi3.Operation{}

// 					input := op.input

// 					path, err := op.route.GetPathTemplate()
// 					method := http.MethodGet
// 					if err != nil {
// 						continue
// 					}
// 					switch meth {
// 					case "PATCH":
// 						method = http.MethodPatch
// 						break
// 					case "POST":
// 						method = http.MethodPost
// 						break
// 					case "DELETE":
// 						method = http.MethodDelete
// 						break
// 					case "PUT":
// 						method = http.MethodPut
// 						break
// 					}

// 					reflector.SetRequest(&oasOp, input, method) // op = GET avec OrderItemRequestParams tant qu'input
// 					//reflector.SetRequest(&oasOp, op.Input[string(meth)], http.MethodGet)      // op = GET avec OrderItemRequestParams tant qu'input
// 					//reflector.SetJSONResponse(&oasOp, op.Output[string(meth)], http.StatusOK) // Si OK (200), l'objet retour est du type OrderItemFilter
// 					// reflector.SetJSONResponse(&op, new([]OrderItemFilter), http.StatusConflict) // Si OK (409), l'objet retour est du type []OrderItemFilter
// 					s.AddOperation(method, path, oasOp)
// 				}
// 			}
// 		}
// 	}

// 	b, err := assertjson.MarshalIndentCompact(s, "", " ", 100) // Générer le resultat tant que doc JSON
// 	if err == nil {
// 		fmt.Print("NO ERROR")
// 	}
// 	ioutil.WriteFile("openapiROUTER.json", b, 0644) // écrire la doc résultat dans un file

// 	return nil
// }

func (r *Router) GroupRoutes() error {
	// "Parse" the routes and read their "name", "method", "query", "path", "body"

	// Groupe the routes by path

	// Return a map "path" - Route that will be read afterward to create the OAS document

	return nil
}
