package catalog

import skill "github.com/mlctrez/lexstream/smapiv1/skill"

/*
LastUpdateRequest Contains attributes related to last modification request of a resource.
*/
type LastUpdateRequest struct {
	Status *CatalogStatusType `json:"status,omitempty"`
	// The version id of the entity returned.
	Version string                     `json:"version,omitempty"`
	Errors  []*skill.StandardizedError `json:"errors,omitempty"`
}

/*
{
 "description": "Contains attributes related to last modification request of a resource.",
 "properties": {
  "errors": {
   "items": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   },
   "type": "array"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.interactionModel.catalog.CatalogStatusType",
   "x-isEnum": true
  },
  "version": {
   "description": "The version id of the entity returned.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
