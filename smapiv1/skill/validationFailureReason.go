package skill

/*
ValidationFailureReason Object representing what is wrong in the request.
*/
type ValidationFailureReason struct {
	// Enum for type of validation failure in the request.
	Type_ *ValidationFailureType `json:"type,omitempty"`
}

/*
{
 "description": "Object representing what is wrong in the request.",
 "properties": {
  "type": {
   "$ref": "#/definitions/v1.skill.ValidationFailureType",
   "description": "Enum for type of validation failure in the request.",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
