package conflictdetection

type ConflictIntentSlot struct {
	Value string `json,omitempty:"value"`
	Type_ string `json,omitempty:"type"`
}

/*
{
 "properties": {
  "type": {
   "type": "string"
  },
  "value": {
   "type": "string"
  }
 },
 "required": [
  "type"
 ],
 "type": "object"
}
*/
