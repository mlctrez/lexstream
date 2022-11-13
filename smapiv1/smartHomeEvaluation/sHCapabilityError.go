package smarthomeevaluation

type SHCapabilityError struct {
	Code    *SHCapabilityErrorCode `json:"code,omitempty"`
	Message string                 `json:"message,omitempty"`
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
