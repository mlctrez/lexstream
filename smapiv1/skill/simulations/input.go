package simulations

type Input struct {
	/*
	   A string corresponding to the utterance text of what a customer would say to Alexa.
	*/
	Content string `json:"content"`
}

/*
{
 "properties": {
  "content": {
   "description": "A string corresponding to the utterance text of what a customer would say to Alexa.\n",
   "type": "string"
  }
 },
 "required": [
  "content"
 ],
 "type": "object"
}
*/
