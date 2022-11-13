package manifest

/*
AlexaForBusinessApis Defines the structure of alexaForBusiness api in the skill manifest.
*/
type AlexaForBusinessApis struct {
	// Contains an array of the supported <region> Objects.
	Regions  map[string]Region      `json:"regions,omitempty"`
	Endpoint *SkillManifestEndpoint `json:"endpoint,omitempty"`
	// Contains the list of supported interfaces.
	Interfaces []*AlexaForBusinessInterface `json:"interfaces,omitempty"`
}

/*
{
 "description": "Defines the structure of alexaForBusiness api in the skill manifest.",
 "properties": {
  "endpoint": {
   "$ref": "#/definitions/v1.skill.Manifest.SkillManifestEndpoint"
  },
  "interfaces": {
   "description": "Contains the list of supported interfaces.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.AlexaForBusinessInterface"
   },
   "type": "array"
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
