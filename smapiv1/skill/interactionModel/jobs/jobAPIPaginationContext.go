package jobs

type JobAPIPaginationContext struct {
	NextToken string `json,omitempty:"nextToken"`
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
