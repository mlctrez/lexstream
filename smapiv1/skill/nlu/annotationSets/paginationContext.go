package annotationsets

type PaginationContext struct {
	// Opaque token returned if there are more results for the given inputs than `maxResults` from the request.
	NextToken string `json:"nextToken,omitempty"`
}

/*
{
 "properties": {
  "nextToken": {
   "description": "Opaque token returned if there are more results for the given inputs than `maxResults` from the request.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
