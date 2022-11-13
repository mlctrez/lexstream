package evaluations

/*
ResolutionsPerAuthorityStatus An object representing the status of entity resolution for the slot.
*/
type ResolutionsPerAuthorityStatus struct {
	Code *ResolutionsPerAuthorityStatusCode `json,omitempty:"code"`
}

/*
{
 "description": "An object representing the status of entity resolution for the slot.",
 "properties": {
  "code": {
   "$ref": "#/definitions/v1.skill.evaluations.ResolutionsPerAuthorityStatusCode",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
