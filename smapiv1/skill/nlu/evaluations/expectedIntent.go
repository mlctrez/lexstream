package evaluations

type ExpectedIntent struct {
	Name  string                              `json,omitempty:"name"`
	Slots map[string]ExpectedIntentSlotsProps `json,omitempty:"slots"`
}

/*
{
 "properties": {
  "name": {
   "type": "string"
  },
  "slots": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.nlu.evaluations.ExpectedIntentSlotsProps"
   },
   "type": "object"
  }
 },
 "type": "object"
}
*/
