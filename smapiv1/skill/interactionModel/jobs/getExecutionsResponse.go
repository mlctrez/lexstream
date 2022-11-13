package jobs

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
GetExecutionsResponse The response of get execution history.
*/
type GetExecutionsResponse struct {
	PaginationContext *JobAPIPaginationContext `json:"paginationContext"`
	Links             *smapiv1.Links           `json:"_links"`
	Executions        []*Execution             `json:"executions"`
}

/*
{
 "description": "The response of get execution history.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "executions": {
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.jobs.Execution"
   },
   "type": "array"
  },
  "paginationContext": {
   "$ref": "#/definitions/v1.skill.interactionModel.jobs.JobAPIPaginationContext"
  }
 },
 "type": "object"
}
*/
