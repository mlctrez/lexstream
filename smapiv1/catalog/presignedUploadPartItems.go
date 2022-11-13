package catalog

import "time"

type PresignedUploadPartItems struct {
	Url        string    `json:"url"`
	PartNumber int       `json:"partNumber"`
	ExpiresAt  time.Time `json:"expiresAt"`
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
