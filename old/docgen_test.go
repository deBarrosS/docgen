package old

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

type OrderItemFilter struct {
	IDs              []string `json:"ids"`
	ItemIDs          []string `json:"item_ids" getorders:"isdefault"`
	LineItemGroupIDs []string `json:"line_item_group_ids" getorders:"isdefault"`
	OrderID          string
}

func (OrderItemFilter) ForceRequestBody() {}

type OrderItemIn struct {
	Id          string                 `path:"in_path"`
	Information map[string]interface{} `json:"information"`
}

func (OrderItemIn) ForceRequestBody() {}

type OrderItemRequestParamsPath struct {
	Id string `path:"id" json:"id"` // IF WE USE THE SAME STRUCT IN DIFFERENT ROUTES PATH IS CONSIDERED OBLIGATORY THEREFORE IT SUCKS :/
	OrderItemRequestParams
}

func (OrderItemRequestParamsPath) ForceRequestBody() {}

type OrderItemRequestParams struct {
	SiteID           string          `json:"site_id" validate:"required,alphanum" create:"required,alphanum"  edit:"required,alphanum"  get:"required,alphanum" addli:"required,alphanum"`
	Fields           []string        `json:"fields"`
	Filter           OrderItemFilter `json:"filter"`
	ItemFeaturesLang *string         `json:"item_features_lang"`
	OrderItem        *OrderItemIn    `json:"order_item" edit:"required"`
}

func (OrderItemRequestParams) ForceRequestBody() {}

func TestCreationApi(t *testing.T) {
	api := new(Api)
	api.InitApi("OMS", "V0", 22)

	r := new(mux.Router)

	api.Add(r.NewRoute().Path("/orders/{id}").Methods("GET")).BodyInput("GET", new(OrderItemRequestParamsPath))
	// api.Add(r.NewRoute().Path("/orders/{id}")).BodyInput("POST", new(OrderItemIn))
	// api.Add(r.NewRoute().Path("/orders/{id2}")).BodyInput("POST", new(OrderItemIn))
	// api.Add(r.NewRoute().Path("/orders/{id3}")).BodyInput("POST", new(OrderItemIn))
	// api.Add(r.NewRoute().Path("/orders/{id4}")).BodyInput("POST", new(OrderItemIn))
	// api.Add(r.NewRoute().Path("/orders/{id5}")).BodyInput("POST", new(OrderItemIn))
	// api.Add(r.NewRoute().Path("/orders/{id6}")).BodyInput("POST", new(OrderItemIn))

	if api != nil {
		fmt.Printf("API Created \n")
		fmt.Printf("Len(api.AllRoutes) = %d \n", len(api.AllRoutes))
		s, err := api.AllRoutes[0].Route.GetPathTemplate()
		inputJson, err := json.MarshalIndent(api.AllRoutes[0].Input, " ", " ")
		if err == nil {
			fmt.Printf("Path Template = %q \n", s)
			fmt.Println("Body Template = \n " + string(inputJson))

		}

	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("paramHandler")
}

func TestDocGen(t *testing.T) {
	// ----------------- API INITIALIZATION ----- Once per service
	api := new(Api)
	api.InitApi("OMS", "V0", 30)
	// ----------------- END OF API INITIALIZATION ----- Once per service

	r := new(mux.Router)

	// ROUTES DECLARATION
	orderPath := "/orders"

	// BEFORE
	r.NewRoute().Path(orderPath).Methods("GET")

	// AFTER
	api.Add(
		r.NewRoute().Path(orderPath).Methods("GET").Handler()).
		BodyInput("GET", new(OrderItemRequestParamsPath)).
		BodyOutput("200", new(OrderItemFilter))

		// TODO :
	// NOTE : Les structs dans les inputs doivent impl??menter l'interface RequestBodyEnforcer de openapi-go (si out of the box)

	api.Add(r.NewRoute().Path(orderPath+"/{id}").Methods("GET")).BodyInput("GET", new(OrderItemRequestParamsPath)) // FINALEMENT PAS BESOIN DE LA METHODE (1 Route par {path+methode})

	api.Add(r.NewRoute().Path(orderPath).Methods("POST")).BodyInput("POST", new(OrderItemRequestParams))

	api.Add(r.NewRoute().Path(orderPath+"/{id}").Methods("PATCH")).BodyInput("PATCH", new(OrderItemRequestParamsPath))

	api.Add(r.NewRoute().Path(orderPath+"/{id}/comments").Methods("POST")).BodyInput("POST", new(OrderItemRequestParamsPath))

	api.Add(r.NewRoute().Path(orderPath + "/{id}/comments").Methods("GET"))

	api.Add(r.NewRoute().Path(orderPath + "/{id}/exclusive_line_items").Methods("PATCH"))

	pathDPDetails := orderPath + "/{id}/delivery_promise/details"

	api.Add(r.NewRoute().Path(pathDPDetails).Methods(http.MethodGet))

	api.Add(r.NewRoute().Path(pathDPDetails).Methods(http.MethodPatch))

	api.Add(r.NewRoute().Path(pathDPDetails).Methods(http.MethodPut))

	api.Add(r.NewRoute().Path(pathDPDetails).Methods(http.MethodDelete))

	// SEARCH_ORDERS

	api.Add(r.NewRoute().Path("/search_orders").Methods("GET"))

	// ORDER ITEMS.

	api.Add(r.NewRoute().Path("/orders/{order_id}/order_items").Methods("GET"))

	api.Add(r.NewRoute().Path("/orders/{order_id}/order_items/{order_item_id}").Methods("PATCH"))

	// LINE ITEM GROUPS.

	api.Add(r.NewRoute().Path("/line_item_groups").Methods("GET"))

	api.Add(r.NewRoute().Path("/line_item_groups/{id}").Methods("GET"))

	api.Add(r.NewRoute().Path("/line_item_groups").Methods("PATCH"))

	api.Add(r.NewRoute().Path("/line_item_groups/{id}").Methods("PATCH"))

	api.Add(r.NewRoute().Path("/line_item_groups/{line_item_group_id}/epcs").Methods("PATCH"))

	api.Add(r.NewRoute().Path("/line_item_groups/{line_item_group_id}/epcs").Methods("DELETE"))

	// PARCELS

	api.Add(r.NewRoute().Path("/parcels").Methods("GET"))

	api.Add(r.NewRoute().Path("/parcels/{id}").Methods("GET"))

	api.Add(r.NewRoute().
		Path("/orders/{order_id}/parcels").
		Methods("POST"))

	api.Add(r.NewRoute().Path("/parcels/{id}").Methods("PATCH"))

	api.Add(r.NewRoute().
		Path("/orders/{order_id}/parcels/{parcel_id}/line_item_groups").
		Methods("POST"))

	// ALERTS

	api.Add(r.NewRoute().Path("/alerts").Methods("GET"))

	// DOCUMENTS

	api.Add(r.NewRoute().Path("/parcels/{id}/documents").Methods("POST"))

	api.Add(r.NewRoute().Path("/parcels/documents").Methods("POST"))

	api.Add(r.NewRoute().Path("/orders/{id}/documents").Methods("POST"))

	// HOMEPAGE

	api.Add(r.NewRoute().Path("/").Methods("GET"))

	// ------- API DOCUMENTATION GENERATION ---- Once per service
	api.GenerateDoc()

}

func TestDocGenHttp(t *testing.T) {

	r := new(mux.Router)

	// ROUTES DECLARATION
	orderPath := "/orders"

	// BEFORE
	r.NewRoute().Path(orderPath).Methods("GET").Handler(http.HandlerFunc(handler))
	orderPath = "only to add breakpoint"
}
