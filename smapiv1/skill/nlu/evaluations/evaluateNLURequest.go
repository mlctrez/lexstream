package evaluations

type EvaluateNLURequest struct {
	Stage  map[string]any `json:"stage"`
	Locale string         `json:"locale"`
	Source *Source        `json:"source"`
}

/*
{
 "properties": {
  "locale": {
   "type": "string"
  },
  "source": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.Source"
  },
  "stage": {
   "properties": {},
   "type": "object"
  }
 },
 "required": [
  "locale",
  "source",
  "stage"
 ],
 "type": "object"
}
*/
