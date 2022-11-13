package manifest

/*
SkillManifestEndpoint Defines the structure for endpoint information in the skill manifest.
*/
type SkillManifestEndpoint struct {
	// Amazon Resource Name (ARN) of the skill's Lambda function or HTTPS URL.
	Uri                string              `json:"uri"`
	SslCertificateType *SSLCertificateType `json:"sslCertificateType"`
}

/*
{
 "description": "Defines the structure for endpoint information in the skill manifest.",
 "properties": {
  "sslCertificateType": {
   "$ref": "#/definitions/v1.skill.Manifest.SSLCertificateType",
   "x-isEnum": true
  },
  "uri": {
   "description": "Amazon Resource Name (ARN) of the skill's Lambda function or HTTPS URL.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
