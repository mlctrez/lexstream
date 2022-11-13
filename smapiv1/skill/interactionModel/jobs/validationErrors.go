package jobs

/*
ValidationErrors The list of errors.
*/
type ValidationErrors struct {
	// The list of errors.
	Errors []*DynamicUpdateError `json,omitempty:"errors"`
}

/*
{
 "description": "The list of errors.",
 "properties": {
  "errors": {
   "description": "The list of errors.",
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.jobs.DynamicUpdateError"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
