package simulations

type Metrics struct {
	/*
	   How long, in milliseconds, it took the skill's Lambda or HTTPS endpoint to process the request.
	*/
	SkillExecutionTimeInMilliseconds int `json:"skillExecutionTimeInMilliseconds"`
}

/*
{
 "properties": {
  "skillExecutionTimeInMilliseconds": {
   "description": "How long, in milliseconds, it took the skill's Lambda or HTTPS endpoint to process the request.\n",
   "type": "integer"
  }
 },
 "type": "object"
}
*/
