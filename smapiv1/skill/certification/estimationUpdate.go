package certification

import "time"

/*
EstimationUpdate Structure for any updates to estimation completion time for certification review for the skill.
*/
type EstimationUpdate struct {
	// Timestamp for originally estimated completion of certification review for the skill.
	OriginalEstimatedCompletionTimestamp time.Time `json:"originalEstimatedCompletionTimestamp"`
	// Timestamp for originally estimated completion of certification review for the skill.
	RevisedEstimatedCompletionTimestamp time.Time `json:"revisedEstimatedCompletionTimestamp"`
	// Reason for updates to estimates for certification review
	Reason string `json:"reason"`
}

/*
{
 "description": "Structure for any updates to estimation completion time for certification review for the skill.",
 "properties": {
  "originalEstimatedCompletionTimestamp": {
   "description": "Timestamp for originally estimated completion of certification review for the skill.",
   "format": "date-time",
   "type": "string"
  },
  "reason": {
   "description": "Reason for updates to estimates for certification review",
   "type": "string"
  },
  "revisedEstimatedCompletionTimestamp": {
   "description": "Timestamp for originally estimated completion of certification review for the skill.",
   "format": "date-time",
   "type": "string"
  }
 },
 "type": "object"
}
*/
