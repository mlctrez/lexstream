package smarthomeevaluation

type TestCaseResult struct {
	ActualCapabilityStates   []*SHCapabilityState   `json:"actualCapabilityStates"`
	ActualResponse           *SHCapabilityResponse  `json:"actualResponse"`
	Directive                *SHCapabilityDirective `json:"directive"`
	Error                    *SHCapabilityError     `json:"error"`
	ExpectedCapabilityStates []*SHCapabilityState   `json:"expectedCapabilityStates"`
	ExpectedResponse         *SHCapabilityResponse  `json:"expectedResponse"`
	Name                     string                 `json:"name"`
	Status                   *TestCaseResultStatus  `json:"status"`
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
