package history

/*
DialogAct The dialog act used in the interaction.
*/
type DialogAct struct {
	Name *DialogActName `json:"name"`
}

/*
{
 "description": "The dialog act used in the interaction.",
 "properties": {
  "name": {
   "$ref": "#/definitions/v1.skill.history.DialogActName",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
