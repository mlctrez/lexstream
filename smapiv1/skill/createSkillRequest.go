package skill

import (
	alexahosted "github.com/mlctrez/lexstream/smapiv1/skill/AlexaHosted"
	manifest "github.com/mlctrez/lexstream/smapiv1/skill/Manifest"
)

type CreateSkillRequest struct {
	// ID of the vendor owning the skill.
	VendorId string                            `json:"vendorId,omitempty"`
	Manifest *manifest.SkillManifest           `json:"manifest,omitempty"`
	Hosting  *alexahosted.HostingConfiguration `json:"hosting,omitempty"`
}

/*
{
 "properties": {
  "hosting": {
   "$ref": "#/definitions/v1.skill.AlexaHosted.HostingConfiguration"
  },
  "manifest": {
   "$ref": "#/definitions/v1.skill.Manifest.SkillManifest"
  },
  "vendorId": {
   "description": "ID of the vendor owning the skill.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
