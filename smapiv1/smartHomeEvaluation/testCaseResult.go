package smarthomeevaluation

type TestCaseResult struct {
	ActualCapabilityStates   []*SHCapabilityState   `json,omitempty:"actualCapabilityStates"`
	ActualResponse           *SHCapabilityResponse  `json,omitempty:"actualResponse"`
	Directive                *SHCapabilityDirective `json,omitempty:"directive"`
	Error                    *SHCapabilityError     `json,omitempty:"error"`
	ExpectedCapabilityStates []*SHCapabilityState   `json,omitempty:"expectedCapabilityStates"`
	ExpectedResponse         *SHCapabilityResponse  `json,omitempty:"expectedResponse"`
	Name                     string                 `json,omitempty:"name"`
	Status                   *TestCaseResultStatus  `json,omitempty:"status"`
}

/*
{
 "properties": {
  "actualCapabilityStates": {
   "items": {
    "$ref": "#/definitions/v1.smartHomeEvaluation.SHCapabilityState"
   },
   "type": "array"
  },
  "actualResponse": {
   "$ref": "#/definitions/v1.smartHomeEvaluation.SHCapabilityResponse"
  },
  "directive": {
   "$ref": "#/definitions/v1.smartHomeEvaluation.SHCapabilityDirective"
  },
  "error": {
   "$ref": "#/definitions/v1.smartHomeEvaluation.SHCapabilityError"
  },
  "expectedCapabilityStates": {
   "items": {
    "$ref": "#/definitions/v1.smartHomeEvaluation.SHCapabilityState"
   },
   "type": "array"
  },
  "expectedResponse": {
   "$ref": "#/definitions/v1.smartHomeEvaluation.SHCapabilityResponse"
  },
  "name": {
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v1.smartHomeEvaluation.TestCaseResultStatus",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
