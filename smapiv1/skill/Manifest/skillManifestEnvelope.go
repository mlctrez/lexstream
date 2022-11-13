package manifest

type SkillManifestEnvelope struct {
	Manifest *SkillManifest `json,omitempty:"manifest"`
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
