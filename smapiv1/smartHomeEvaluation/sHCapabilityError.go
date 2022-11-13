package smarthomeevaluation

type SHCapabilityError struct {
	Code    *SHCapabilityErrorCode `json,omitempty:"code"`
	Message string                 `json,omitempty:"message"`
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
