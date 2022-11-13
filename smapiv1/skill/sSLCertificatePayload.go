package skill

type SSLCertificatePayload struct {
	// The default ssl certificate for the skill. If a request is made for a region without an explicit ssl certificate, this certificate will be used.
	SslCertificate string `json:"sslCertificate,omitempty"`
	// A map of region to ssl certificate. Keys are string region codes (https://developer.amazon.com/docs/smapi/skill-manifest.html#regions), values are regional ssl certificate objects which contain the ssl certificate blobs as strings.
	Regions map[string]RegionalSSLCertificate `json:"regions,omitempty"`
}

/*
{
 "properties": {
  "regions": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.RegionalSSLCertificate"
   },
   "description": "A map of region to ssl certificate. Keys are string region codes (https://developer.amazon.com/docs/smapi/skill-manifest.html#regions), values are regional ssl certificate objects which contain the ssl certificate blobs as strings.",
   "type": "object"
  },
  "sslCertificate": {
   "description": "The default ssl certificate for the skill. If a request is made for a region without an explicit ssl certificate, this certificate will be used.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
