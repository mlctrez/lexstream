package evaluations

type EvaluateNLURequest struct {
	Stage  map[string]any `json:"stage,omitempty"`
	Locale string         `json:"locale,omitempty"`
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
 "required": [
  "locale",
  "source",
  "stage"
 ],
 "type": "object"
}
*/
