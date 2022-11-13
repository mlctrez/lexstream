package skill

/*
ResourceImportStatus Defines the structure for a resource deployment status.
*/
type ResourceImportStatus struct {
	// Resource name. eg. manifest, interactionModels.en_US and so on.
	Name     string               `json,omitempty:"name"`
	Status   *ResponseStatus      `json,omitempty:"status"`
	Action   *Action              `json,omitempty:"action"`
	Errors   []*StandardizedError `json,omitempty:"errors"`
	Warnings []*StandardizedError `json,omitempty:"warnings"`
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
