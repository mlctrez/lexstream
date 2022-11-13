package skill

/*
ExportResponseSkill Defines the structure of the GetExport response.
*/
type ExportResponseSkill struct {
	ETag     string `json:"eTag,omitempty"`
	Location string `json:"location,omitempty"`
	// ExpiresAt timestamp in milliseconds.
	ExpiresAt string `json:"expiresAt,omitempty"`
}

/*
{
 "description": "Defines the structure of the GetExport response.",
 "properties": {
  "eTag": {
   "type": "string"
  },
  "expiresAt": {
   "description": "ExpiresAt timestamp in milliseconds.",
   "type": "string"
  },
  "location": {
   "format": "uri",
   "type": "string"
  }
 },
 "type": "object"
}
*/
