package skill

import "time"

/*
UploadResponse Defines the structure for skill upload response.
*/
type UploadResponse struct {
	// The upload URL to upload a skill package.
	UploadUrl string    `json:"uploadUrl,omitempty"`
	ExpiresAt time.Time `json:"expiresAt,omitempty"`
}

/*
{
 "description": "Defines the structure for skill upload response.",
 "properties": {
  "expiresAt": {
   "format": "date-time",
   "type": "string"
  },
  "uploadUrl": {
   "description": "The upload URL to upload a skill package.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
