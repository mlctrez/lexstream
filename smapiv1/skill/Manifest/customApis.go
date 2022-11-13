package manifest

/*
CustomApis Defines the structure for custom api of the skill.
*/
type CustomApis struct {
	// Contains an array of the supported <region> Objects.
	Regions  map[string]Region      `json,omitempty:"regions"`
	Endpoint *SkillManifestEndpoint `json,omitempty:"endpoint"`
	// Defines the structure for interfaces supported by the custom api of the skill.
	Interfaces []*Interface `json,omitempty:"interfaces"`
	// List of provided tasks.
	Tasks       []*SkillManifestCustomTask `json,omitempty:"tasks"`
	Connections *CustomConnections         `json,omitempty:"connections"`
}

/*
{
 "description": "Defines the structure for custom api of the skill.",
 "properties": {
  "connections": {
   "$ref": "#/definitions/v1.skill.Manifest.CustomConnections"
  },
  "endpoint": {
   "$ref": "#/definitions/v1.skill.Manifest.SkillManifestEndpoint"
  },
  "interfaces": {
   "description": "Defines the structure for interfaces supported by the custom api of the skill.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.Interface"
   },
   "type": "array"
  },
  "regions": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.Manifest.Region"
   },
   "description": "Contains an array of the supported \u003cregion\u003e Objects.",
   "type": "object"
  },
  "tasks": {
   "description": "List of provided tasks.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.SkillManifestCustomTask"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
