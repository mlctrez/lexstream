package manifest

type MusicRequest struct {
	// Name of the request.
	Name string `json:"name,omitempty"`
}

/*
{
 "properties": {
  "name": {
   "description": "Name of the request.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
