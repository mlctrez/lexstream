package evaluations

type Results struct {
	/*
	   url to get the test case result details Example: /v1/skills/{skillId}/nluEvaluations/{evaluationId}/results
	*/
	Href string `json:"href,omitempty"`
}

/*
{
 "properties": {
  "href": {
   "description": "url to get the test case result details Example: /v1/skills/{skillId}/nluEvaluations/{evaluationId}/results\n",
   "type": "string"
  }
 },
 "type": "object"
}
*/
