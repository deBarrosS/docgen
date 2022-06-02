package docgen

import (
	"github.com/gorilla/mux"
)

type Router struct {
	routes       []Route
	nbRoutes     int
	defaultResps map[int]interface{}
}

type Route struct {
	path    string
	method  string
	input   interface{}         // No need to map by method because of the line above ^^
	resps   map[int]interface{} // outputs per return code
	handler func()
}

func InitApi(name, version string) *Router {
	r := new(Router)
	r.routes = make([]Route, 30)
	r.nbRoutes = 0
	// r.doc = new(Api)
	// r.doc.InitApi(name, version, 30)

	return r
}

func NewRouter() *Router {
	r := new(Router)
	//r.routes = make([]Route, 30)
	r.nbRoutes = 0
	return r
}
func (r *Router) NewRoute(path, meth string, handler func()) *Router {

	r.routes = append(r.routes, Route{
		path:    path,
		method:  meth,
		handler: handler,
	})

	return r
}

// Router as param or as caller ?
// Pass "app" as parameter -> Get the real router and add it to the handlers
func RegisterRoutes(r *Router) (*mux.Router, error) {
	router := new(mux.Router)

	for _, route := range r.routes {
		router.NewRoute().Path(route.path).Methods(route.method) //Add handler
		// Shall we add some error handling here? If an error, then log (the x route couldn't be set and continue ?)
	}

	return router, nil
}
