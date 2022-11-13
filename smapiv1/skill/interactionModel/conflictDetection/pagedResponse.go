package conflictdetection

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

type PagedResponse struct {
	PaginationContext *PaginationContext `json:"paginationContext,omitempty"`
	Links             *smapiv1.Links     `json:"_links,omitempty"`
}

/*
{
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "paginationContext": {
   "$ref": "#/definitions/v1.skill.interactionModel.conflictDetection.PaginationContext"
  }
 },
 "type": "object"
}
*/
