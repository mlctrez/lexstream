package skill

/*
StandardizedError Standardized structure which wraps machine parsable and human readable information about an error.
*/
type StandardizedError struct {
	// Standardized, machine readable structure that wraps all the information about a specific occurrence of an error of the type specified by the code.
	ValidationDetails *ValidationDetails `json:"validationDetails"`
}

/*
{
 "description": "Standardized structure which wraps machine parsable and human readable information about an error.",
 "properties": {
  "validationDetails": {
   "$ref": "#/definitions/v1.skill.validationDetails",
   "description": "Standardized, machine readable structure that wraps all the information about a specific occurrence of an error of the type specified by the code."
  }
 },
 "type": "object",
 "x-inheritsFrom": "v1.Error"
}
*/
