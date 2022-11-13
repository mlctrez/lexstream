package upload

import smapiv0 "github.com/mlctrez/lexstream/smapiv0"

type ListUploadsResponse struct {
	Links       *smapiv0.Links `json:"_links"`
	IsTruncated bool           `json:"isTruncated"`
	NextToken   string         `json:"nextToken"`
	// List of upload summaries.
	Uploads []*ContentUploadSummary `json:"uploads"`
}

/*
{
 "properties": {
  "_links": {
   "$ref": "#/definitions/v0.Links"
  },
  "isTruncated": {
   "type": "boolean"
  },
  "nextToken": {
   "type": "string"
  },
  "uploads": {
   "description": "List of upload summaries.",
   "items": {
    "$ref": "#/definitions/v0.catalog.upload.ContentUploadSummary"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
