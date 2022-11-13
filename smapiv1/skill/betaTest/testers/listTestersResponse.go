package testers

type ListTestersResponse struct {
	Testers     []*TesterWithDetails `json:"testers"`
	IsTruncated bool                 `json:"isTruncated"`
	NextToken   string               `json:"nextToken"`
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
