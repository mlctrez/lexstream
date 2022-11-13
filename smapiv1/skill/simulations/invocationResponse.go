package simulations

type InvocationResponse struct {
	/*
	   Payload that was returned by the skill's Lambda or HTTPS endpoint.
	*/
	Body map[string]any `json,omitempty:"body"`
}

/*
{
 "properties": {
  "body": {
   "additionalProperties": {
    "properties": {},
    "type": "object"
   },
   "description": "Payload that was returned by the skill's Lambda or HTTPS endpoint.\n",
   "type": "object"
  }
 },
 "type": "object"
}
*/
