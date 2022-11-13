package evaluations

type EvaluationInputs struct {
	Locale string         `json:"locale"`
	Stage  map[string]any `json:"stage"`
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
 "type": "object"
}
*/
