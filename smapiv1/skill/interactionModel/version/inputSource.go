package version

/*
InputSource Definition for catalog version input data.
*/
type InputSource struct {
	// Type of catalog.
	Type_ string `json:"type,omitempty"`
	// Url to the catalog reference.
	Url string `json:"url,omitempty"`
}

/*
{
 "description": "Definition for catalog version input data.",
 "properties": {
  "type": {
   "description": "Type of catalog.",
   "type": "string"
  },
  "url": {
   "description": "Url to the catalog reference.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
