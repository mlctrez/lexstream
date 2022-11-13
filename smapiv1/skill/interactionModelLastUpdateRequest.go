package skill

/*
InteractionModelLastUpdateRequest Contains attributes related to last modification (create/update) request of a resource.
*/
type InteractionModelLastUpdateRequest struct {
	Status       *Status              `json,omitempty:"status"`
	Errors       []*StandardizedError `json,omitempty:"errors"`
	Warnings     []*StandardizedError `json,omitempty:"warnings"`
	BuildDetails *BuildDetails        `json,omitempty:"buildDetails"`
}

/*
{
 "description": "Contains attributes related to last modification (create/update) request of a resource.",
 "properties": {
  "buildDetails": {
   "$ref": "#/definitions/v1.skill.BuildDetails"
  },
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
