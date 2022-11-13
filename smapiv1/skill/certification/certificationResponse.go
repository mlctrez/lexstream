package certification

import "time"

type CertificationResponse struct {
	// Certification Id for the skill
	Id     string               `json,omitempty:"id"`
	Status *CertificationStatus `json,omitempty:"status"`
	// Timestamp for when the skill was submitted for certification.
	SkillSubmissionTimestamp time.Time            `json,omitempty:"skillSubmissionTimestamp"`
	ReviewTrackingInfo       *ReviewTrackingInfo  `json,omitempty:"reviewTrackingInfo"`
	Result                   *CertificationResult `json,omitempty:"result"`
}

/*
{
 "properties": {
  "id": {
   "description": "Certification Id for the skill",
   "type": "string"
  },
  "result": {
   "$ref": "#/definitions/v1.skill.certification.CertificationResult"
  },
  "reviewTrackingInfo": {
   "$ref": "#/definitions/v1.skill.certification.ReviewTrackingInfo"
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
