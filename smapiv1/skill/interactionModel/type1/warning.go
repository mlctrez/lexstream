package type1

/*
Warning The warning which would not fail requests.
*/
type Warning struct {
	// The warning code.
	Code string `json:"code"`
	// The warning message.
	Message string `json:"message"`
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
