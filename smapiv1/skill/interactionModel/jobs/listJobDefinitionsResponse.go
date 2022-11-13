package jobs

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
ListJobDefinitionsResponse The response of list job definitions.
*/
type ListJobDefinitionsResponse struct {
	PaginationContext *JobAPIPaginationContext `json:"paginationContext"`
	Links             *smapiv1.Links           `json:"_links"`
	Jobs              []*JobDefinitionMetadata `json:"jobs"`
}

/*
{
 "description": "The response of list job definitions.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "jobs": {
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.jobs.JobDefinitionMetadata"
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
