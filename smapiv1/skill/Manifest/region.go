package manifest

/*
Region Defines the structure for regional information.
*/
type Region struct {
	Endpoint *SkillManifestEndpoint `json,omitempty:"endpoint"`
}

/*
{
 "description": "Defines the structure for regional information.",
 "properties": {
  "endpoint": {
   "$ref": "#/definitions/v1.skill.Manifest.SkillManifestEndpoint"
  }
 },
 "type": "object"
}
*/
