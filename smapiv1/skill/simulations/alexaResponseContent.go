package simulations

type AlexaResponseContent struct {
	// The plain text get from Alexa speech response
	Caption string `json:"caption,omitempty"`
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
