package skill

import "time"

/*
RollbackRequestStatus Rollback request for a skill
*/
type RollbackRequestStatus struct {
	// rollback request id
	Id string `json,omitempty:"id"`
	// The target skill version to rollback to.
	TargetVersion  string                      `json,omitempty:"targetVersion"`
	SubmissionTime time.Time                   `json,omitempty:"submissionTime"`
	Status         *RollbackRequestStatusTypes `json,omitempty:"status"`
	Errors         []*StandardizedError        `json,omitempty:"errors"`
}

/*
{
 "description": "Rollback request for a skill",
 "properties": {
  "errors": {
   "items": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   },
   "type": "array"
  },
  "id": {
   "description": "rollback request id",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.RollbackRequestStatusTypes",
   "x-isEnum": true
  },
  "submissionTime": {
   "format": "date-time",
   "type": "string"
  },
  "targetVersion": {
   "description": "The target skill version to rollback to.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
