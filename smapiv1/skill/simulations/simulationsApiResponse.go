package simulations

type SimulationsApiResponse struct {
	// Id of the simulation resource.
	Id     string                        `json,omitempty:"id"`
	Status *SimulationsApiResponseStatus `json,omitempty:"status"`
	Result *SimulationResult             `json,omitempty:"result"`
}

/*
{
 "properties": {
  "id": {
   "description": "Id of the simulation resource.",
   "type": "string"
  },
  "result": {
   "$ref": "#/definitions/v1.skill.simulations.SimulationResult"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.simulations.SimulationsApiResponseStatus",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
