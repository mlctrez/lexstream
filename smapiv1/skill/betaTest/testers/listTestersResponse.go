package testers

type ListTestersResponse struct {
	Testers     []*TesterWithDetails `json:"testers,omitempty"`
	IsTruncated bool                 `json:"isTruncated,omitempty"`
	NextToken   string               `json:"nextToken,omitempty"`
}

/*
{
 "properties": {
  "isTruncated": {
   "type": "boolean"
  },
  "nextToken": {
   "type": "string"
  },
  "testers": {
   "items": {
    "$ref": "#/definitions/v1.skill.betaTest.testers.TesterWithDetails"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
