package evaluations

type PagedResultsResponsePaginationContext struct {
	// opaque token returned if there are more results for the given inputs than `maxResults` from the request.
	NextToken string `json:"nextToken,omitempty"`
	// Total available results for the given query.
	TotalCount string `json:"totalCount,omitempty"`
}

/*
{
 "properties": {
  "nextToken": {
   "description": "opaque token returned if there are more results for the given inputs than `maxResults` from the request.",
   "type": "string"
  },
  "totalCount": {
   "description": "Total available results for the given query.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
