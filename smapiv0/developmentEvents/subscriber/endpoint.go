package subscriber

type Endpoint struct {
	// Uri of the endpoint that receives the notification.
	Uri           string                 `json:"uri"`
	Authorization *EndpointAuthorization `json:"authorization"`
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
