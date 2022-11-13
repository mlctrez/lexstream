package invocations

type Metrics struct {
	/*
	   How long, in milliseconds, it took the skill's Lambda or HTTPS endpoint to process the request.
	*/
	SkillExecutionTimeInMilliseconds int `json,omitempty:"skillExecutionTimeInMilliseconds"`
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
