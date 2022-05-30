package docgen

import (
	"errors"
	"fmt"

	"github.com/gorilla/mux"
)

type Api struct {
	AllRoutes []*Operation
}

type Operation struct {
	Route  *mux.Route
	Input  map[string]interface{}
	Output map[string]interface{}
}

// As we are returning pointers, no problem regarding the changes to the struct
func (api *Api) Add(newRoute *mux.Route) (*Operation, error) {
	if api == nil {
		api = new(Api)
	}

	op := new(Operation)
	op.Route = newRoute

	if api.AllRoutes == nil {
		return nil, errors.New("Api not initialized")
	}

	nRoutes := len(api.AllRoutes)
	api.AllRoutes[nRoutes+1] = op
	fmt.Sprint("Route added")
	return op, nil
}

func (op *Operation) BodyInput(meth string, body interface{}) (*Operation, error) {
	if op == nil {
		return nil, errors.New("Operation not initialized")
	}
	if op.Input == nil {
		op.Input = make(map[string]interface{})
	}
	op.Input[meth] = body

	return op, nil
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

// func AddOperations(route *mux.Route) error {
// 	//Get the path of the route and the methods it matches against
// 	path, err := route.GetPathRegexp()
// 	if err != nil {
// 		return errors.New("Error getting the Path.")
// 	}
// 	methods, err := route.GetMethods()

// 	if err != nil {
// 		return errors.New("Error getting the Methods.")
// 	}

// 	for _, meth := range methods {
// 		if Doc.operations[path] == nil {
// 			Operations[path] = meth
// 			continue
// 		}
// 		Doc.Operations[path] = Doc.Operations[path] + meth
// 	}

// 	return nil
// }

// func GenDoc(router *mux.Router) {

// }
// func Hello(name string) (string, error) {
// 	if name == "" {
// 		return "", errors.New("empty name")
// 	}

// 	//message := fmt.Sprintf(randomFormat(), name)
// 	message := fmt.Sprintf(randomFormat())
// 	return message, nil
// }

// func Hellos(names []string) (map[string]string, error) {
// 	messages := make(map[string]string)
// 	for _, name := range names {
// 		message, err := Hello(name)
// 		if err != nil {
// 			return nil, err
// 		}
// 		messages[name] = message
// 	}
// 	return messages, nil
// }

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

// func randomFormat() string {
// 	formats := []string{
// 		"Hi, %v. Welcome!",
// 		"Great to see you, %v!",
// 		"Hail, %v ! Well met!",
// 	}
// 	return formats[rand.Intn(len(formats))]
// }
