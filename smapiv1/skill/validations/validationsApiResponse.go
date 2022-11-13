package validations

type ValidationsApiResponse struct {
	// Id of the validation resource.
	Id     string                        `json:"id,omitempty"`
	Status *ValidationsApiResponseStatus `json:"status,omitempty"`
	Result *ValidationsApiResponseResult `json:"result,omitempty"`
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
