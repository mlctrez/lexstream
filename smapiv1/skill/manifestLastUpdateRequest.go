package skill

/*
ManifestLastUpdateRequest Contains attributes related to last modification (create/update) request of a resource.
*/
type ManifestLastUpdateRequest struct {
	Status   *Status              `json:"status"`
	Errors   []*StandardizedError `json:"errors"`
	Warnings []*StandardizedError `json:"warnings"`
	// on success, this field indicates the created version.
	Version string `json:"version"`
}

/*
{
 "description": "Contains attributes related to last modification (create/update) request of a resource.",
 "properties": {
  "errors": {
   "items": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   },
   "type": "array"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.Status",
   "x-isEnum": true
  },
  "version": {
   "description": "on success, this field indicates the created version.",
   "type": "string"
  },
  "warnings": {
   "items": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
