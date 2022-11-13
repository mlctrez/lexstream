package evaluations

type EvaluationItems struct {
	EvaluationMetadata
	// evaluation id
	Id string `json,omitempty:"id"`
}

/*
{
 "properties": {
  "id": {
   "description": "evaluation id",
   "type": "string"
  }
 },
 "required": [
  "id"
 ],
 "type": "object"
}
*/
