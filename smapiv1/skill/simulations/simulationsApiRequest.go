package simulations

type SimulationsApiRequest struct {
	Input   *Input   `json,omitempty:"input"`
	Device  *Device  `json,omitempty:"device"`
	Session *Session `json,omitempty:"session"`
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
