package smarthomeevaluation

type SHCapabilityError struct {
	Code    *SHCapabilityErrorCode `json:"code"`
	Message string                 `json:"message"`
}

/*
{
 "properties": {
  "code": {
   "$ref": "#/definitions/v1.smartHomeEvaluation.SHCapabilityErrorCode",
   "x-isEnum": true
  },
  "message": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
