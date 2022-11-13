package conflictdetection

type PaginationContext struct {
	// A token returned if there are more results for the given inputs than `maxResults` from the request. It should also be used in the next request to retrieve more results.
	NextToken string `json:"nextToken,omitempty"`
	// Total avaliable results for the given query.
	TotalCount int `json:"totalCount,omitempty"`
}

/*
{
 "properties": {
  "nextToken": {
   "description": "A token returned if there are more results for the given inputs than `maxResults` from the request. It should also be used in the next request to retrieve more results.",
   "type": "string"
  },
  "totalCount": {
   "description": "Total avaliable results for the given query.",
   "format": "int64",
   "type": "integer"
  }
 },
 "type": "object"
}
*/
