package manifest

/*
SmartHomeApis Defines the structure for smart home api of the skill.
*/
type SmartHomeApis struct {
	// Contains an array of the supported <region> Objects.
	Regions         map[string]LambdaRegion `json,omitempty:"regions"`
	Endpoint        *LambdaEndpoint         `json,omitempty:"endpoint"`
	ProtocolVersion *SmartHomeProtocol      `json,omitempty:"protocolVersion"`
}

/*
{
 "description": "Defines the structure for smart home api of the skill.",
 "properties": {
  "endpoint": {
   "$ref": "#/definitions/v1.skill.Manifest.LambdaEndpoint"
  },
  "protocolVersion": {
   "$ref": "#/definitions/v1.skill.Manifest.SmartHomeProtocol",
   "x-isEnum": true
  },
  "regions": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.Manifest.LambdaRegion"
   },
   "description": "Contains an array of the supported \u003cregion\u003e Objects.",
   "type": "object"
  }
 },
 "type": "object"
}
*/
