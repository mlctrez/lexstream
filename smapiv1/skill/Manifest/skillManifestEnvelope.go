package manifest

type SkillManifestEnvelope struct {
	Manifest *SkillManifest `json:"manifest"`
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
