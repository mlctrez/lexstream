package simulations

type AlexaResponseContent struct {
	// The plain text get from Alexa speech response
	Caption string `json,omitempty:"caption"`
}

/*
{
 "properties": {
  "caption": {
   "description": "The plain text get from Alexa speech response",
   "type": "string"
  }
 },
 "type": "object"
}
*/
