package conflictdetection

type ConflictResult struct {
	// Sample utterance provided by 3P developers for intents.
	SampleUtterance string          `json,omitempty:"sampleUtterance"`
	Intent          *ConflictIntent `json,omitempty:"intent"`
}

/*
{
 "properties": {
  "intent": {
   "$ref": "#/definitions/v1.skill.interactionModel.conflictDetection.ConflictIntent"
  },
  "sampleUtterance": {
   "description": "Sample utterance provided by 3P developers for intents.",
   "type": "string"
  }
 },
 "required": [
  "intent",
  "sampleUtterance"
 ],
 "type": "object"
}
*/
