package certification

import "time"

/*
ReviewTrackingInfoSummary Structure for summarised view of review tracking information of the skill. This does not have the estimationUpdates array field.
*/
type ReviewTrackingInfoSummary struct {
	// Timestamp for estimated completion of certification review for the skill.
	EstimatedCompletionTimestamp time.Time `json:"estimatedCompletionTimestamp,omitempty"`
	// Timestamp for actual completion of certification review workflow for the skill.
	ActualCompletionTimestamp time.Time `json:"actualCompletionTimestamp,omitempty"`
	// Timestamp for when the last update was made to review tracking info.
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
}

/*
{
 "description": "Structure for summarised view of review tracking information of the skill. This does not have the estimationUpdates array field.",
 "properties": {
  "actualCompletionTimestamp": {
   "description": "Timestamp for actual completion of certification review workflow for the skill.",
   "format": "date-time",
   "type": "string"
  },
  "estimatedCompletionTimestamp": {
   "description": "Timestamp for estimated completion of certification review for the skill.",
   "format": "date-time",
   "type": "string"
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
