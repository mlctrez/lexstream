package manifest

type HealthInterface struct {
	// Name of the interface.
	Namespace string   `json,omitempty:"namespace"`
	Version   *Version `json,omitempty:"version"`
	// Defines the details of requests that a health skill is capable of handling.
	Requests []*HealthRequest `json,omitempty:"requests"`
	// Defines the list for health skill locale specific publishing information in the skill manifest.
	Locales map[string]LocalizedHealthInfo `json,omitempty:"locales"`
}

/*
{
 "properties": {
  "locales": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.Manifest.LocalizedHealthInfo"
   },
   "description": "Defines the list for health skill locale specific publishing information in the skill manifest.",
   "type": "object"
  },
  "namespace": {
   "description": "Name of the interface.",
   "type": "string"
  },
  "requests": {
   "description": "Defines the details of requests that a health skill is capable of handling.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.HealthRequest"
   },
   "type": "array"
  },
  "version": {
   "$ref": "#/definitions/v1.skill.Manifest.Version",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
