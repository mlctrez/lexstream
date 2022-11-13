package skill

/*
SkillInteractionModelStatus Defines the structure for interaction model build status.
*/
type SkillInteractionModelStatus struct {
	LastUpdateRequest *InteractionModelLastUpdateRequest `json:"lastUpdateRequest"`
	// An opaque identifier for last successfully updated resource.
	ETag string `json:"eTag"`
	// Version for last successfully built model.
	Version string `json:"version"`
}

/*
{
 "description": "Defines the structure for interaction model build status.",
 "properties": {
  "eTag": {
   "description": "An opaque identifier for last successfully updated resource.",
   "type": "string"
  },
  "lastUpdateRequest": {
   "$ref": "#/definitions/v1.skill.InteractionModelLastUpdateRequest"
  },
  "version": {
   "description": "Version for last successfully built model.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
