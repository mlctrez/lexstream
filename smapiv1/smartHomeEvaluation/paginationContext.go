package smarthomeevaluation

type PaginationContext struct {
	NextToken  string `json:"nextToken,omitempty"`
	TotalCount int    `json:"totalCount,omitempty"`
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
