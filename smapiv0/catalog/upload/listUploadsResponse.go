package upload

import smapiv0 "github.com/mlctrez/lexstream/smapiv0"

type ListUploadsResponse struct {
	Links       *smapiv0.Links `json:"_links,omitempty"`
	IsTruncated bool           `json:"isTruncated,omitempty"`
	NextToken   string         `json:"nextToken,omitempty"`
	// List of upload summaries.
	Uploads []*ContentUploadSummary `json:"uploads,omitempty"`
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
