package evaluations

type Evaluation struct {
	EvaluationEntity
	// id of the job
	Id string `json:"id,omitempty"`
}

/*
{
 "properties": {
  "id": {
   "description": "id of the job",
   "type": "string"
  }
 },
 "type": "object"
}
*/
