package validations

type ValidationsApiResponse struct {
	// Id of the validation resource.
	Id     string                        `json:"id"`
	Status *ValidationsApiResponseStatus `json:"status"`
	Result *ValidationsApiResponseResult `json:"result"`
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
