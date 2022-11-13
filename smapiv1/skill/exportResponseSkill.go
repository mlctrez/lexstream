package skill

/*
ExportResponseSkill Defines the structure of the GetExport response.
*/
type ExportResponseSkill struct {
	ETag     string `json,omitempty:"eTag"`
	Location string `json,omitempty:"location"`
	// ExpiresAt timestamp in milliseconds.
	ExpiresAt string `json,omitempty:"expiresAt"`
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
