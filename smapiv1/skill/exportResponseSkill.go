package skill

/*
ExportResponseSkill Defines the structure of the GetExport response.
*/
type ExportResponseSkill struct {
	ETag     string `json:"eTag"`
	Location string `json:"location"`
	// ExpiresAt timestamp in milliseconds.
	ExpiresAt string `json:"expiresAt"`
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
