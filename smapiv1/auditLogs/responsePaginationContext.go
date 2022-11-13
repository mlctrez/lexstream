package auditlogs

/*
ResponsePaginationContext This object contains the next token used to load the next page of the result.
*/
type ResponsePaginationContext struct {
	// This token can be used to load the next page of the result.
	NextToken string `json,omitempty:"nextToken"`
}

/*
{
 "description": "This object contains the next token used to load the next page of the result.",
 "properties": {
  "nextToken": {
   "description": "This token can be used to load the next page of the result.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
