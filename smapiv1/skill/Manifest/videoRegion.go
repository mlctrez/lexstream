package manifest

/*
VideoRegion Defines the structure for endpoint information.
*/
type VideoRegion struct {
	Endpoint *LambdaEndpoint `json,omitempty:"endpoint"`
	// The channel through which the partner skill can communicate to Alexa.
	Upchannel []*UpChannelItems `json,omitempty:"upchannel"`
}

/*
{
 "description": "Defines the structure for endpoint information.",
 "properties": {
  "endpoint": {
   "$ref": "#/definitions/v1.skill.Manifest.LambdaEndpoint"
  },
  "upchannel": {
   "description": "The channel through which the partner skill can communicate to Alexa.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.UpChannelItems"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
