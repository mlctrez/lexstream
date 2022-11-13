package type1

/*
Warning The warning which would not fail requests.
*/
type Warning struct {
	// The warning code.
	Code string `json,omitempty:"code"`
	// The warning message.
	Message string `json,omitempty:"message"`
}

/*
{
 "description": "The warning which would not fail requests.",
 "properties": {
  "code": {
   "description": "The warning code.",
   "type": "string"
  },
  "message": {
   "description": "The warning message.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
