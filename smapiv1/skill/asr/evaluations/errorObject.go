package evaluations

/*
ErrorObject Object containing information about the error occurred during an evaluation run. This filed would present if an unexpected error occurred during an evaluatin run.
*/
type ErrorObject struct {
	// human-readable error message
	Message string `json:"message"`
	// machine-readable error code
	Code string `json:"code"`
}

/*
{
 "description": "Object containing information about the error occurred during an evaluation run. This filed would present if an unexpected error occurred during an evaluatin run.\n",
 "properties": {
  "code": {
   "description": "machine-readable error code",
   "type": "string"
  },
  "message": {
   "description": "human-readable error message",
   "type": "string"
  }
 },
 "required": [
  "code",
  "message"
 ],
 "type": "object"
}
*/
