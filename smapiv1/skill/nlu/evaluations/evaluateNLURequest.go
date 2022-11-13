package evaluations

type EvaluateNLURequest struct {
	Stage  map[string]any `json,omitempty:"stage"`
	Locale string         `json,omitempty:"locale"`
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
 "required": [
  "locale",
  "source",
  "stage"
 ],
 "type": "object"
}
*/
