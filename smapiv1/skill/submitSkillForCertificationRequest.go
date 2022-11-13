package skill

type SubmitSkillForCertificationRequest struct {
	PublicationMethod *PublicationMethod `json:"publicationMethod,omitempty"`
	// Description of the version (limited to 300 characters).
	VersionMessage string `json:"versionMessage,omitempty"`
}

/*
{
 "properties": {
  "publicationMethod": {
   "$ref": "#/definitions/v1.skill.PublicationMethod",
   "x-isEnum": true
  },
  "versionMessage": {
   "description": "Description of the version (limited to 300 characters).",
   "type": "string"
  }
 },
 "type": "object"
}
*/
