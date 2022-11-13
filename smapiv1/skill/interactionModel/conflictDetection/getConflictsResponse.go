package conflictdetection

type GetConflictsResponse struct {
	PagedResponse
	Results []*GetConflictsResponseResult `json,omitempty:"results"`
}

/*
{
 "properties": {
  "results": {
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.conflictDetection.GetConflictsResponseResult"
   },
   "type": "array"
  }
 },
 "required": [
  "results"
 ],
 "type": "object"
}
*/
