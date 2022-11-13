package invocations

type Request struct {
	// Skill's Lambda or HTTPS endpoint.
	Endpoint string `json:"endpoint"`
	/*
	   JSON payload that was sent to the skill's Lambda or HTTPS endpoint.
	*/
	Body map[string]any `json:"body"`
}

/*
{
 "properties": {
  "body": {
   "description": "JSON payload that was sent to the skill's Lambda or HTTPS endpoint.\n",
   "properties": {},
   "type": "object"
  },
  "endpoint": {
   "description": "Skill's Lambda or HTTPS endpoint.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
