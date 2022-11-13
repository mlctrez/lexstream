package evaluations

type Evaluation struct {
	EvaluationEntity
	// id of the job
	Id string `json,omitempty:"id"`
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
