package simulations

type AlexaExecutionInfo struct {
	AlexaResponses []*AlexaResponse `json:"alexaResponses,omitempty"`
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
