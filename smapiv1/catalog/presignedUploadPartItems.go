package catalog

import "time"

type PresignedUploadPartItems struct {
	Url        string    `json,omitempty:"url"`
	PartNumber int       `json,omitempty:"partNumber"`
	ExpiresAt  time.Time `json,omitempty:"expiresAt"`
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
