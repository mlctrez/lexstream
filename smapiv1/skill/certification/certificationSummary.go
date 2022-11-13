package certification

import "time"

/*
CertificationSummary Summary of the certification resource. This is a leaner view of the certification resource for the collections API.
*/
type CertificationSummary struct {
	// Certification Id for the skill.
	Id     string               `json:"id,omitempty"`
	Status *CertificationStatus `json:"status,omitempty"`
	// Timestamp for when the skill was submitted for certification.
	SkillSubmissionTimestamp time.Time                  `json:"skillSubmissionTimestamp,omitempty"`
	ReviewTrackingInfo       *ReviewTrackingInfoSummary `json:"reviewTrackingInfo,omitempty"`
}

/*
{
 "description": "Summary of the certification resource. This is a leaner view of the certification resource for the collections API.",
 "properties": {
  "id": {
   "description": "Certification Id for the skill.",
   "type": "string"
  },
  "reviewTrackingInfo": {
   "$ref": "#/definitions/v1.skill.certification.ReviewTrackingInfoSummary"
  },
  "skillSubmissionTimestamp": {
   "description": "Timestamp for when the skill was submitted for certification.",
   "format": "date-time",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.certification.CertificationStatus",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
