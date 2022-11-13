package skill

/*
ManifestStatus Defines the structure for a resource status.
*/
type ManifestStatus struct {
	LastUpdateRequest *ManifestLastUpdateRequest `json,omitempty:"lastUpdateRequest"`
	// An opaque identifier for last successfully updated resource.
	ETag string `json,omitempty:"eTag"`
}

/*
{
 "description": "Defines the structure for a resource status.",
 "properties": {
  "eTag": {
   "description": "An opaque identifier for last successfully updated resource.",
   "type": "string"
  },
  "lastUpdateRequest": {
   "$ref": "#/definitions/v1.skill.ManifestLastUpdateRequest"
  }
 },
 "type": "object"
}
*/
