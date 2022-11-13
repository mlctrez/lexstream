package manifest

/*
HealthApis Defines the structure of health api in the skill manifest.
*/
type HealthApis struct {
	// Contains an array of the supported <region> Objects.
	Regions         map[string]Region      `json,omitempty:"regions"`
	Endpoint        *SkillManifestEndpoint `json,omitempty:"endpoint"`
	ProtocolVersion *HealthProtocolVersion `json,omitempty:"protocolVersion"`
	Interfaces      *HealthInterface       `json,omitempty:"interfaces"`
}

/*
{
 "description": "Defines the structure of health api in the skill manifest.",
 "properties": {
  "endpoint": {
   "$ref": "#/definitions/v1.skill.Manifest.SkillManifestEndpoint"
  },
  "interfaces": {
   "$ref": "#/definitions/v1.skill.Manifest.HealthInterface"
  },
  "protocolVersion": {
   "$ref": "#/definitions/v1.skill.Manifest.HealthProtocolVersion",
   "x-isEnum": true
  },
  "regions": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.Manifest.Region"
   },
   "description": "Contains an array of the supported \u003cregion\u003e Objects.",
   "type": "object"
  }
 },
 "type": "object"
}
*/
