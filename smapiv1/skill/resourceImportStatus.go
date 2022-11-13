package skill

/*
ResourceImportStatus Defines the structure for a resource deployment status.
*/
type ResourceImportStatus struct {
	// Resource name. eg. manifest, interactionModels.en_US and so on.
	Name     string               `json:"name,omitempty"`
	Status   *ResponseStatus      `json:"status,omitempty"`
	Action   *Action              `json:"action,omitempty"`
	Errors   []*StandardizedError `json:"errors,omitempty"`
	Warnings []*StandardizedError `json:"warnings,omitempty"`
}

/*
{
 "description": "Defines the structure for a resource deployment status.",
 "properties": {
  "action": {
   "$ref": "#/definitions/v1.skill.Action",
   "x-isEnum": true
  },
  "errors": {
   "items": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   },
   "type": "array"
  },
  "name": {
   "description": "Resource name. eg. manifest, interactionModels.en_US and so on.",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.ResponseStatus",
   "x-isEnum": true
  },
  "warnings": {
   "items": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   },
   "type": "array"
  }
 },
 "required": [
  "status"
 ],
 "type": "object"
}
*/
