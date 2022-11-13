package isp

/*
ListInSkillProductResponse List of in-skill product response.
*/
type ListInSkillProductResponse struct {
	InSkillProductSummaryList *ListInSkillProduct `json:"inSkillProductSummaryList"`
}

/*
{
 "description": "List of in-skill product response.",
 "properties": {
  "inSkillProductSummaryList": {
   "$ref": "#/definitions/v1.isp.ListInSkillProduct"
  }
 },
 "type": "object"
}
*/
