package catalog

/*
CatalogStatus Defines the structure for catalog status response.
*/
type CatalogStatus struct {
	LastUpdateRequest *LastUpdateRequest `json:"lastUpdateRequest,omitempty"`
}

/*
{
 "description": "Defines the structure for catalog status response.",
 "properties": {
  "lastUpdateRequest": {
   "$ref": "#/definitions/v1.skill.interactionModel.catalog.LastUpdateRequest"
  }
 },
 "type": "object"
}
*/
