package manifest

/*
SkillManifestEvents Defines the structure for subscribed events information in the skill manifest.
*/
type SkillManifestEvents struct {
	// Contains an array of eventName object each of which contains the name of a skill event.
	Subscriptions []*EventName         `json:"subscriptions"`
	Publications  []*EventPublications `json:"publications"`
	// Contains an array of the supported <region> Objects.
	Regions  map[string]Region      `json:"regions"`
	Endpoint *SkillManifestEndpoint `json:"endpoint"`
}

/*
{
 "description": "Defines the structure for subscribed events information in the skill manifest.",
 "properties": {
  "endpoint": {
   "$ref": "#/definitions/v1.skill.Manifest.SkillManifestEndpoint"
  },
  "publications": {
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.EventPublications"
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
  "subscriptions": {
   "description": "Contains an array of eventName object each of which contains the name of a skill event.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.EventName"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
