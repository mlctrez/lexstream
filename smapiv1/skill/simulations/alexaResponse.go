package simulations

type AlexaResponse struct {
	// The type of Alexa response
	Type_ string `json,omitempty:"type"`
	/*
	   The detail information needs to exposed in this type of Alexa response.
	*/
	Content *AlexaResponseContent `json,omitempty:"content"`
}

/*
{
 "properties": {
  "content": {
   "$ref": "#/definitions/v1.skill.simulations.AlexaResponseContent",
   "description": "The detail information needs to exposed in this type of Alexa response.\n"
  },
  "type": {
   "description": "The type of Alexa response",
   "type": "string"
  }
 },
 "type": "object"
}
*/
