package skill

import (
	alexahosted "github.com/mlctrez/lexstream/smapiv1/skill/AlexaHosted"
	manifest "github.com/mlctrez/lexstream/smapiv1/skill/Manifest"
)

type CreateSkillRequest struct {
	// ID of the vendor owning the skill.
	VendorId string                            `json,omitempty:"vendorId"`
	Manifest *manifest.SkillManifest           `json,omitempty:"manifest"`
	Hosting  *alexahosted.HostingConfiguration `json,omitempty:"hosting"`
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
