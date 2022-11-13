package simulations

/*
Session Session settings for running current simulation.
*/
type Session struct {
	Mode *SessionMode `json:"mode"`
}

/*
{
 "description": "Session settings for running current simulation.\n",
 "properties": {
  "mode": {
   "$ref": "#/definitions/v1.skill.simulations.SessionMode",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
