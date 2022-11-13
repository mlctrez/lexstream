package smarthomeevaluation

type PaginationContext struct {
	NextToken  string `json:"nextToken"`
	TotalCount int    `json:"totalCount"`
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
