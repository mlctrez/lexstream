package jobs

type JobAPIPaginationContext struct {
	NextToken string `json:"nextToken,omitempty"`
}

/*
{
 "properties": {
  "nextToken": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
