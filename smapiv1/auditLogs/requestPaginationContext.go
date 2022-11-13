package auditlogs

/*
RequestPaginationContext This object includes nextToken and maxResults.
*/
type RequestPaginationContext struct {
	// When the response to this API call is truncated, the response includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that this API understands. Token has expiry of 1 hour.
	NextToken string `json,omitempty:"nextToken"`
	// Sets the maximum number of results returned in the response body. If you want to retrieve more or less than the default of 50 results, you can add this parameter to your request. maxResults can exceed the upper limit of 250 but we will not return more items than that. The response might contain fewer results than maxResults for purpose of keeping SLA or because there are not enough items, but it will never contain more.
	MaxResults int `json,omitempty:"maxResults"`
}

/*
{
 "description": "This object includes nextToken and maxResults.",
 "properties": {
  "maxResults": {
   "description": "Sets the maximum number of results returned in the response body. If you want to retrieve more or less than the default of 50 results, you can add this parameter to your request. maxResults can exceed the upper limit of 250 but we will not return more items than that. The response might contain fewer results than maxResults for purpose of keeping SLA or because there are not enough items, but it will never contain more.",
   "maximum": 250,
   "minimum": 1,
   "multipleOf": 1,
   "type": "number"
  },
  "nextToken": {
   "description": "When the response to this API call is truncated, the response includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that this API understands. Token has expiry of 1 hour.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
