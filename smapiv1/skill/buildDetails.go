package skill

/*
BuildDetails Contains array which describes steps involved in a build. Elements (or build steps) are added
to this array as they become IN_PROGRESS.
*/
type BuildDetails struct {
	// An array where each element represents a build step.
	Steps []*BuildStep `json,omitempty:"steps"`
}

/*
{
 "description": "Contains array which describes steps involved in a build. Elements (or build steps) are added\nto this array as they become IN_PROGRESS.\n",
 "properties": {
  "steps": {
   "description": "An array where each element represents a build step.",
   "items": {
    "$ref": "#/definitions/v1.skill.BuildStep"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
