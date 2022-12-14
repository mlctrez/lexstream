package skill

/*
WithdrawRequest The payload for the withdraw operation.
*/
type WithdrawRequest struct {
	Reason *Reason `json:"reason,omitempty"`
	// The message only in case the reason in OTHER.
	Message string `json:"message,omitempty"`
}

/*
{
 "description": "The payload for the withdraw operation.",
 "properties": {
  "message": {
   "description": "The message only in case the reason in OTHER.",
   "type": "string"
  },
  "reason": {
   "$ref": "#/definitions/v1.skill.Reason",
   "x-isEnum": true
  }
 },
 "required": [
  "reason"
 ],
 "type": "object"
}
*/
