package validations

type ValidationsApiResponse struct {
	// Id of the validation resource.
	Id     string                        `json,omitempty:"id"`
	Status *ValidationsApiResponseStatus `json,omitempty:"status"`
	Result *ValidationsApiResponseResult `json,omitempty:"result"`
}

/*
{
 "properties": {
  "id": {
   "description": "Id of the validation resource.",
   "type": "string"
  },
  "result": {
   "$ref": "#/definitions/v1.skill.validations.ValidationsApiResponseResult"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.validations.ValidationsApiResponseStatus",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
