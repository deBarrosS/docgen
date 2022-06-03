package docgen

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) // TODO :  add di.Container

func (HTTPHandler) ServeHTTP(http.ResponseWriter, *http.Request) {

}

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
	handler HTTPHandler
}

func InitApi(name, version string) *Router {
	r := new(Router)
	r.routes = make([]Route, 30)
	r.nbRoutes = 0

	return r
}

func NewRouter() *Router {
	r := new(Router)
	r.nbRoutes = 0
	return r
}
func (r *Router) NewRoute(path, meth string) *Router {

	r.routes = append(r.routes, Route{
		path:   path,
		method: meth,
	})

	return r
}

func CreateHTTPHandler(h HTTPHandler, input interface{}, out200 interface{}) {
	// Among other stuff, return the function like we already do, but this time, calling the handler function passing the parameters needed
}

// Router as param or as caller ? -> Param for keeping the calling pattern
// Pass "app" as parameter -> Get the real router and add it to the handlers
func RegisterRoutes(r *Router) (*mux.Router, error) {
	router := new(mux.Router)

	for _, route := range r.routes {
		if &route != nil && &route.path != nil && &route.method != nil {
			router.NewRoute().Path(route.path).Methods(route.method).Handler(route.handler) //Add handler
		}
		// Shall we add some error handling here? If an error, then log (the x route couldn't be set and continue ?)
	}

	return router, nil
}
