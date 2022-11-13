package conflictdetection

type GetConflictsResponseResult struct {
	// Utterance resolved from sample utterance that causes conflicts among different intents.
	ConflictingUtterance string            `json:"conflictingUtterance,omitempty"`
	Conflicts            []*ConflictResult `json:"conflicts,omitempty"`
}

/*
{
 "properties": {
  "conflictingUtterance": {
   "description": "Utterance resolved from sample utterance that causes conflicts among different intents.",
   "type": "string"
  },
  "conflicts": {
   "items": {
    "$ref": "#/definitions/v1.skill.interactionModel.conflictDetection.ConflictResult"
   },
   "type": "array"
  }
 },
 "required": [
  "conflictingUtterance",
  "conflicts"
 ],
 "type": "object"
}
*/
