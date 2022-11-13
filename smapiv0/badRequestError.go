package smapiv0

type BadRequestError struct {
	// Human readable description of error.
	Message string `json,omitempty:"message"`
	// An array of violation messages.
	Violations []*Error `json,omitempty:"violations"`
}

/*
{
 "properties": {
  "message": {
   "description": "Human readable description of error.",
   "type": "string"
  },
  "violations": {
   "description": "An array of violation messages.",
   "items": {
    "$ref": "#/definitions/v0.Error"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
