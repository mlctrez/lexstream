package smapiv0

type Error struct {
	/*
	   Error code that maps to an error message. Developers with different locales should be able to lookup the error description based on this code.
	*/
	Code string `json,omitempty:"code"`
	// Readable description of error.
	Message string `json,omitempty:"message"`
}

/*
{
 "properties": {
  "code": {
   "description": "Error code that maps to an error message. Developers with different locales should be able to lookup the error description based on this code.\n",
   "type": "string"
  },
  "message": {
   "description": "Readable description of error.",
   "type": "string"
  }
 },
 "required": [
  "message"
 ],
 "type": "object"
}
*/
