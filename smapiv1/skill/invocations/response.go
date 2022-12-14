package invocations

type Response struct {
	/*
	   Payload that was returned by the skill's Lambda or HTTPS endpoint.
	*/
	Body map[string]any `json:"body,omitempty"`
}

/*
{
 "properties": {
  "body": {
   "description": "Payload that was returned by the skill's Lambda or HTTPS endpoint.\n",
   "properties": {},
   "type": "object"
  }
 },
 "type": "object"
}
*/
