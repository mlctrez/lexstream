package manifest

/*
LambdaEndpoint Contains the uri field. This sets the global default endpoint.
*/
type LambdaEndpoint struct {
	// Amazon Resource Name (ARN) of the skill's Lambda function or HTTPS URL.
	Uri string `json:"uri"`
}

/*
{
 "description": "Contains the uri field. This sets the global default endpoint.",
 "properties": {
  "uri": {
   "description": "Amazon Resource Name (ARN) of the skill's Lambda function or HTTPS URL.",
   "type": "string"
  }
 },
 "required": [
  "uri"
 ],
 "type": "object"
}
*/
