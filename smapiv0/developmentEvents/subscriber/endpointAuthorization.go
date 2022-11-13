package subscriber

/*
EndpointAuthorization Authorization information to be able to publish notification to specified endpoint.
*/
type EndpointAuthorization struct {
	Type_ string `json:"type,omitempty"`
}

/*
{
 "description": "Authorization information to be able to publish notification to specified endpoint.",
 "discriminator": "type",
 "properties": {
  "type": {
   "type": "string",
   "x-isDiscriminator": true
  }
 },
 "required": [
  "type"
 ],
 "type": "object"
}
*/
