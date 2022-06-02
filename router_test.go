package docgen

import (
	"testing"
)

func TestCreationRouter(t *testing.T) {
	//------------ Initialization of the router variable

	// BEFORE
	// r := app.Get("router").(*mux.Router)

	// AFTER
	// router := app.Get("router").(*mux.Router) // CHANGE BECAUSE OF env variables
	// router := new(mux.Router)
	// r := new(Router)
	// r.router = router

	// ------------ END OF INITIALIZATION

	// Nothing to change afterwards

}

func TestDocGenRouter(t *testing.T) {

	// ----- INITIALIZATION

	// r	 := InitApi("OMS", "V0")
	// router := new(mux.Router)
	// r.router = router

	// // ----- END OF INITIALIZATION

	// // ----- ROUTES DECLARATION (as normal (or almost))
	// orderPath := "/orders"

	// r.NewRoute().Path(orderPath).Methods("GET")//.Handler(func inputn output)

	// // ------- API DOCUMENTATION GENERATION
	// r.GenerateDoc()

}
