package evaluations

type Metrics struct {
	// overall error rate for the ASR evaluation run
	OverallErrorRate int `json:"overallErrorRate"`
}

/*
{
 "properties": {
  "overallErrorRate": {
   "description": "overall error rate for the ASR evaluation run",
   "type": "number"
  }
 },
 "required": [
  "overallErrorRate"
 ],
 "type": "object"
}
*/
