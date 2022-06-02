package docgen

import (
	"fmt"
	"testing"
)

func TestNewRouter(t *testing.T) {
	r := NewRouter()
	r.routes = []Route{
		{
			path:   "/",
			method: "get",
		},
	}

	fmt.Println("r.routes[0].path")
	fmt.Println(r.routes[0].path)
	fmt.Println(r.routes[0].method)
}
func TestRegisterRoutes(t *testing.T) {

}

func TestDocgenRouter(t *testing.T) {
	r := NewRouter() // TODO: add the di.Container as parameter
	r.routes = []Route{
		{
			path:   "/",
			method: "get",
		},
		{
			path:   "/",
			method: "post",
		},
		{
			path:   "/",
			method: "patch",
		},
		{
			path:   "/withbody",
			method: "post",
		},
		{
			path:   "/withbody",
			method: "patch",
		},
	}

	// Server and Documentation use the same structure but are not mutually dependent
	RegisterRoutes(r)
	GenerateDoc(r)
}
