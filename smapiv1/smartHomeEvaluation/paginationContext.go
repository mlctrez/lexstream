package smarthomeevaluation

type PaginationContext struct {
	NextToken  string `json,omitempty:"nextToken"`
	TotalCount int    `json,omitempty:"totalCount"`
}

/*
{
 "properties": {
  "nextToken": {
   "type": "string"
  },
  "totalCount": {
   "format": "int64",
   "type": "integer"
  }
 },
 "type": "object"
}
*/
