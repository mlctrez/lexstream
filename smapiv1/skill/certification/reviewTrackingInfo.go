package certification

import "time"

/*
ReviewTrackingInfo Structure for review tracking information of the skill.
*/
type ReviewTrackingInfo struct {
	// Timestamp for estimated completion of certification review for the skill.
	EstimatedCompletionTimestamp time.Time `json:"estimatedCompletionTimestamp"`
	// Timestamp for actual completion of certification review for the skill.
	ActualCompletionTimestamp time.Time `json:"actualCompletionTimestamp"`
	// Timestamp for when the last update was made to review tracking info.
	LastUpdated time.Time `json:"lastUpdated"`
	// List of updates to estimation completion time for certification review for the skill.
	EstimationUpdates []*EstimationUpdate `json:"estimationUpdates"`
}

/*
{
 "description": "Structure for review tracking information of the skill.",
 "properties": {
  "actualCompletionTimestamp": {
   "description": "Timestamp for actual completion of certification review for the skill.",
   "format": "date-time",
   "type": "string"
  },
  "estimatedCompletionTimestamp": {
   "description": "Timestamp for estimated completion of certification review for the skill.",
   "format": "date-time",
   "type": "string"
  },
  "estimationUpdates": {
   "description": "List of updates to estimation completion time for certification review for the skill.",
   "items": {
    "$ref": "#/definitions/v1.skill.certification.EstimationUpdate"
   },
   "type": "array"
  },
  "lastUpdated": {
   "description": "Timestamp for when the last update was made to review tracking info.",
   "format": "date-time",
   "type": "string"
  }
 },
 "type": "object"
}
*/
