package evaluations

type EvaluationItems struct {
	EvaluationMetadata
	// evaluation id
	Id string `json:"id,omitempty"`
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
