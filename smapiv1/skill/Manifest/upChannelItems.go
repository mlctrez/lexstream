package manifest

type UpChannelItems struct {
	// Use \"SNS\" for this field.
	Type_ string `json:"type"`
	// SNS Amazon Resource Name (ARN) for video skill through which video partner can send events to Alexa.
	Uri string `json:"uri"`
}

/*
{
 "properties": {
  "type": {
   "description": "Use \\\"SNS\\\" for this field.",
   "type": "string"
  },
  "uri": {
   "description": "SNS Amazon Resource Name (ARN) for video skill through which video partner can send events to Alexa.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
