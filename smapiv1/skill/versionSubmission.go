package skill

import "time"

/*
VersionSubmission Submission for a skill version
*/
type VersionSubmission struct {
	Status         *VersionSubmissionStatus `json:"status"`
	SubmissionTime time.Time                `json:"submissionTime"`
}

/*
{
 "description": "Submission for a skill version\n",
 "properties": {
  "status": {
   "$ref": "#/definitions/v1.skill.VersionSubmissionStatus",
   "x-isEnum": true
  },
  "submissionTime": {
   "format": "date-time",
   "type": "string"
  }
 },
 "type": "object"
}
*/