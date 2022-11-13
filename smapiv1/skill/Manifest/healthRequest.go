package manifest

type HealthRequest struct {
	// Defines the name of request, each request has their own payload format.
	Name string `json:"name,omitempty"`
}

/*
{
 "properties": {
  "name": {
   "description": "Defines the name of request, each request has their own payload format.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
