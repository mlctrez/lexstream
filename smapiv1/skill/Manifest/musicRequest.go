package manifest

type MusicRequest struct {
	// Name of the request.
	Name string `json,omitempty:"name"`
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
