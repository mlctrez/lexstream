package subscriber

type Endpoint struct {
	// Uri of the endpoint that receives the notification.
	Uri           string                 `json:"uri,omitempty"`
	Authorization *EndpointAuthorization `json:"authorization,omitempty"`
}

/*
{
 "properties": {
  "authorization": {
   "$ref": "#/definitions/v0.developmentEvents.subscriber.EndpointAuthorization"
  },
  "uri": {
   "description": "Uri of the endpoint that receives the notification.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
