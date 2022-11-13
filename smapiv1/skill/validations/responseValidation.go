package validations

type ResponseValidation struct {
	/*
	   Short, human readable title of the validation performed.
	*/
	Title string `json,omitempty:"title"`
	/*
	   Human readable description of the validation performed. May include instructions to address validation failure.
	*/
	Description string `json,omitempty:"description"`
	/*
	   Dot-delimited category.
	*/
	Category string `json,omitempty:"category"`
	/*
	   Locale of the validation.
	*/
	Locale     string                        `json,omitempty:"locale"`
	Importance *ResponseValidationImportance `json,omitempty:"importance"`
	Status     *ResponseValidationStatus     `json,omitempty:"status"`
}

/*
{
 "properties": {
  "category": {
   "description": "Dot-delimited category.\n",
   "type": "string"
  },
  "description": {
   "description": "Human readable description of the validation performed. May include instructions to address validation failure.\n",
   "type": "string"
  },
  "importance": {
   "$ref": "#/definitions/v1.skill.validations.ResponseValidationImportance",
   "x-isEnum": true
  },
  "locale": {
   "description": "Locale of the validation.\n",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.validations.ResponseValidationStatus",
   "x-isEnum": true
  },
  "title": {
   "description": "Short, human readable title of the validation performed.\n",
   "type": "string"
  }
 },
 "type": "object"
}
*/
