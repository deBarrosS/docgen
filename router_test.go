package docgen

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/deBarrosS/docgen/old"
)

func TestNewRouter(t *testing.T) {
	r := NewRouter()
	r.routes = []Route{
		{
			path:   "/",
			method: "gEt",
		},
	}

	fmt.Println("r.routes[0].path")
	fmt.Println(r.routes[0].path)
	fmt.Println(r.routes[0].method)
}
func TestRegisterRoutes(t *testing.T) {

}

func TestDocgen(t *testing.T) {
	r := NewRouter() // TODO: add the di.Container as parameter
	r.routes = []Route{
		{
			path:   "/",
			method: "gEt",
		},
		{
			path:   "/",
			method: http.MethodDelete,
		},
		{
			path:   "/",
			method: "patch",
		},
		{
			path:   "/withbody",
			method: "POST",
		},
		{
			path:   "/withbody",
			method: "gEt",
		},
	}

	// Server and Documentation use the same structure but are not mutually dependent
	muxrouter, err := RegisterRoutes(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(muxrouter) // no errors linting

	GenerateDoc(r)
}

func TestDocgenResponses(t *testing.T) {
	r := NewRouter() // TODO: add the di.Container as parameter
	r.routes = []Route{
		{
			path:   "/{name}/{id}",
			method: "post",
			input:  new(old.OrderItemFilter),
			resps: map[int]interface{}{
				200:                   new(old.OrderItemFilter),
				201:                   new(old.OrderItemFilter),
				http.StatusBadRequest: new(old.OtherJson),
			},
		},
		{
			path:   "/{id}",
			method: http.MethodDelete,
		},
		{
			path:   "/{names}",
			method: "patch",
		},
		{
			path:   "/withbody",
			method: "POST",
		},
		{
			path:   "/withbody",
			method: "gEt",
		},
	}

	// Server and Documentation use the same structure but are not mutually dependent
	muxrouter, err := RegisterRoutes(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(muxrouter) // no errors linting

	GenerateDoc(r)
}
