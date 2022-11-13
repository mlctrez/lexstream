package jobs

type JobAPIPaginationContext struct {
	NextToken string `json:"nextToken"`
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
