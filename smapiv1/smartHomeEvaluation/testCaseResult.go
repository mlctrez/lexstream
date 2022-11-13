package smarthomeevaluation

type TestCaseResult struct {
	ActualCapabilityStates   []*SHCapabilityState   `json:"actualCapabilityStates,omitempty"`
	ActualResponse           *SHCapabilityResponse  `json:"actualResponse,omitempty"`
	Directive                *SHCapabilityDirective `json:"directive,omitempty"`
	Error                    *SHCapabilityError     `json:"error,omitempty"`
	ExpectedCapabilityStates []*SHCapabilityState   `json:"expectedCapabilityStates,omitempty"`
	ExpectedResponse         *SHCapabilityResponse  `json:"expectedResponse,omitempty"`
	Name                     string                 `json:"name,omitempty"`
	Status                   *TestCaseResultStatus  `json:"status,omitempty"`
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
