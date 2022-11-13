package testers

type ListTestersResponse struct {
	Testers     []*TesterWithDetails `json,omitempty:"testers"`
	IsTruncated bool                 `json,omitempty:"isTruncated"`
	NextToken   string               `json,omitempty:"nextToken"`
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
