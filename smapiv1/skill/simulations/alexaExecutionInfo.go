package simulations

type AlexaExecutionInfo struct {
	AlexaResponses []*AlexaResponse `json,omitempty:"alexaResponses"`
}

/*
{
 "properties": {
  "alexaResponses": {
   "items": {
    "$ref": "#/definitions/v1.skill.simulations.AlexaResponse"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
