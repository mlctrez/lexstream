package manifest

type HealthRequest struct {
	// Defines the name of request, each request has their own payload format.
	Name string `json,omitempty:"name"`
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
