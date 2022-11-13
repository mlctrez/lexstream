package evaluations

/*
PaginationContext This holds all data needed to control pagination from the user.
*/
type PaginationContext struct {
	/*
	   The page token, this should be passed as a `nextToken` query parameter to the API to retrieve more items. If this field is not present the end of all of the items was reached. If a `maxResults` query parameter was specified then no more than `maxResults` items are returned.
	*/
	NextToken string `json:"nextToken,omitempty"`
}

/*
{
 "description": "This holds all data needed to control pagination from the user.\n",
 "properties": {
  "nextToken": {
   "description": "The page token, this should be passed as a `nextToken` query parameter to the API to retrieve more items. If this field is not present the end of all of the items was reached. If a `maxResults` query parameter was specified then no more than `maxResults` items are returned.\n",
   "type": "string"
  }
 },
 "required": [
  "nextToken"
 ],
 "type": "object"
}
*/
