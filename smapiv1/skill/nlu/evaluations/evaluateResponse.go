package evaluations

type EvaluateResponse struct {
	// Id used to retrieve the evaluation later.
	Id string `json:"id,omitempty"`
}

/*
{
 "properties": {
  "id": {
   "description": "Id used to retrieve the evaluation later.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
