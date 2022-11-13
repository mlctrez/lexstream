package isp

/*
InSkillProductSummaryResponse In-skill product summary response.
*/
type InSkillProductSummaryResponse struct {
	InSkillProductSummary *InSkillProductSummary `json:"inSkillProductSummary,omitempty"`
}

/*
{
 "description": "In-skill product summary response.",
 "properties": {
  "inSkillProductSummary": {
   "$ref": "#/definitions/v1.isp.InSkillProductSummary"
  }
 },
 "type": "object"
}
*/
