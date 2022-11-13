package evaluations

type PostAsrEvaluationsResponseObject struct {
	// ID used to retrieve the evaluation status/results later.
	Id string `json:"id,omitempty"`
}

/*
{
 "properties": {
  "id": {
   "description": "ID used to retrieve the evaluation status/results later.",
   "type": "string"
  }
 },
 "required": [
  "id"
 ],
 "type": "object"
}
*/
