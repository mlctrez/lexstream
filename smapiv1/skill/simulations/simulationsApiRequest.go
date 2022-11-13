package simulations

type SimulationsApiRequest struct {
	Input   *Input   `json:"input"`
	Device  *Device  `json:"device"`
	Session *Session `json:"session"`
}

/*
{
 "properties": {
  "device": {
   "$ref": "#/definitions/v1.skill.simulations.Device"
  },
  "input": {
   "$ref": "#/definitions/v1.skill.simulations.Input"
  },
  "session": {
   "$ref": "#/definitions/v1.skill.simulations.Session"
  }
 },
 "required": [
  "device",
  "input"
 ],
 "type": "object"
}
*/
