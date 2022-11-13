package evaluations

type EvaluationInputs struct {
	Locale string         `json:"locale,omitempty"`
	Stage  map[string]any `json:"stage,omitempty"`
	Source *Source        `json:"source,omitempty"`
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
