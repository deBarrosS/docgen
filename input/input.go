package input

// Input is an empty structure whose goal is only to implement ForceRequestBody() method
// By doing so, the swaggest/openapi-go allows documenting a GET request containing a body
// How to use : Insert Input on the declaration of the structure you want to use as a GET request's body
// Example :
// type BodyGetRequest struct {
// 	input.Input
// 	BodyProperty  string `json:"body_property"`
// }
type Input struct{}

func (*Input) ForceRequestBody() {}
