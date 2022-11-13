package validations

type ResponseValidation struct {
	/*
	   Short, human readable title of the validation performed.
	*/
	Title string `json:"title,omitempty"`
	/*
	   Human readable description of the validation performed. May include instructions to address validation failure.
	*/
	Description string `json:"description,omitempty"`
	/*
	   Dot-delimited category.
	*/
	Category string `json:"category,omitempty"`
	/*
	   Locale of the validation.
	*/
	Locale     string                        `json:"locale,omitempty"`
	Importance *ResponseValidationImportance `json:"importance,omitempty"`
	Status     *ResponseValidationStatus     `json:"status,omitempty"`
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
