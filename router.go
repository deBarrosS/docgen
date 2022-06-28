package docgen

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) // TODO :  add di.Container

func (HTTPHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
}

type Info struct {
	Name    string
	Version string
}

type Router struct {
	Routes       []*Route
	BaseUrl      string
	DefaultResps map[int]interface{}
}

type Route struct {
	Path       string
	Method     string
	Input      interface{}         // No need to map by method because of the line above ^^
	Resps      map[int]interface{} // outputs per return code
	Handler    http.HandlerFunc
	Internal   bool
	Deprecated bool
}

type Handler struct {
	Function http.HandlerFunc
	Input    interface{}         // No need to map by method because of the line above ^^
	Resps    map[int]interface{} // outputs per return code
}

func InitApi(name, version string) *Router {
	r := new(Router)
	r.Routes = make([]*Route, 2)

	return r
}

func NewRouter() *Router {
	r := new(Router)
	return r
}

func (r *Router) NewRoute() *Route {
	route := new(Route)
	r.Routes = append(r.Routes, route)
	return route
}

func (r *Route) WithPath(path string) *Route { // Naming problems for Path as method and Prop
	r.Path = path
	return r
}

func (r *Route) WithHandler(handler http.HandlerFunc) *Route { // Naming problems for Path as method and Prop
	r.Handler = handler
	return r
}

func (r *Route) WithMethod(method string) *Route { // Naming problems for Path as method and Prop
	r.Method = method
	return r
}

func (r *Route) WithInput(input interface{}) *Route { // Naming problems for Path as method and Prop
	r.Input = input
	return r
}

func (r *Route) WithOutput(responses map[int]interface{}) *Route { // Naming problems for Path as method and Prop
	r.Resps = responses
	return r
}

func CreateHTTPHandler(h HTTPHandler, input interface{}, out200 interface{}) {
	// Among other stuff, return the function like we already do, but this time, calling the handler function passing the parameters needed
}

// Router as param or as caller ? -> Param to keep the function calling pattern
// Pass "app" as parameter -> Get the real router and add it to the handlers
func RegisterRoutes(r *Router) (*mux.Router, error) {
	router := new(mux.Router)

	for _, route := range r.Routes {
		if route != nil && &(route.Path) != nil && &(route.Method) != nil {
			if &(route.Handler) != nil {
				router.NewRoute().Path(route.Path).Methods(route.Method).Handler(route.Handler) //Add handler
			}
		}
		// Shall we add some error handling here? If an error, then log (the x route couldn't be set and continue ?)
		// Or even the route doesn't contain the path or the method or the handler ...
	}

	return router, nil
}

// Router as param or as caller ? -> Param to keep the function calling pattern
// Pass "app" as parameter -> Get the real router and add it to the handlers
func AttachRoutes(r *Router, muxRouter *mux.Router) (*mux.Router, error) {
	for _, route := range r.Routes {
		if &route != nil && &route.Path != nil && &route.Method != nil {
			if &route.Handler != nil {
				muxRouter.NewRoute().Path(route.Path).Methods(route.Method).Handler(route.Handler) //Add handler
			}
		}
		// Shall we add some error handling here? If an error, then log (the x route couldn't be set and continue ?)
		// Or even the route doesn't contain the path or the method or the handler ...
	}
	return muxRouter, nil
}
