package manifest

type SkillManifestEnvelope struct {
	Manifest *SkillManifest `json:"manifest,omitempty"`
}

/*
{
 "properties": {
  "manifest": {
   "$ref": "#/definitions/v1.skill.Manifest.SkillManifest"
  }
 },
 "type": "object"
}
*/
