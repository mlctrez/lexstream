package evaluations

type PaginationContext struct {
	// opaque token returned if there are more results for the given inputs than `maxResults` from the request.
	NextToken string `json,omitempty:"nextToken"`
}

/*
{
 "properties": {
  "nextToken": {
   "description": "opaque token returned if there are more results for the given inputs than `maxResults` from the request.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
