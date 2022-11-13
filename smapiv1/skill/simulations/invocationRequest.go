package simulations

type InvocationRequest struct {
	// Skill's Lambda or HTTPS endpoint.
	Endpoint string `json,omitempty:"endpoint"`
	/*
	   JSON payload that was sent to the skill's Lambda or HTTPS endpoint.
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
   "description": "JSON payload that was sent to the skill's Lambda or HTTPS endpoint.\n",
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
