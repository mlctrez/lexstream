package evaluations

type EvaluationInputs struct {
	Locale string         `json,omitempty:"locale"`
	Stage  map[string]any `json,omitempty:"stage"`
	Source *Source        `json,omitempty:"source"`
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
