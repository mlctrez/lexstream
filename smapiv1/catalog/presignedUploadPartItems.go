package catalog

import "time"

type PresignedUploadPartItems struct {
	Url        string    `json:"url,omitempty"`
	PartNumber int       `json:"partNumber,omitempty"`
	ExpiresAt  time.Time `json:"expiresAt,omitempty"`
}

/*
{
 "properties": {
  "expiresAt": {
   "format": "date-time",
   "type": "string"
  },
  "partNumber": {
   "type": "integer"
  },
  "url": {
   "type": "string"
  }
 },
 "required": [
  "expiresAt",
  "partNumber",
  "url"
 ],
 "type": "object"
}
*/
