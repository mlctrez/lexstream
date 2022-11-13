package validations

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

type ValidationsApiResponseResult struct {
	Validations []*ResponseValidation `json:"validations,omitempty"`
	Error       *smapiv1.Error        `json:"error,omitempty"`
}

/*
{
 "properties": {
  "error": {
   "$ref": "#/definitions/v1.Error"
  },
  "validations": {
   "items": {
    "$ref": "#/definitions/v1.skill.validations.ResponseValidation"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
