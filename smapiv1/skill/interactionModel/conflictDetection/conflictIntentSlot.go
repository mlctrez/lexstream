package conflictdetection

type ConflictIntentSlot struct {
	Value string `json:"value,omitempty"`
	Type_ string `json:"type,omitempty"`
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
