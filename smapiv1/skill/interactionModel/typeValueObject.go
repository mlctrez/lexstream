package interactionmodel

/*
TypeValueObject The object that contains individual type values.
*/
type TypeValueObject struct {
	Value    string   `json:"value,omitempty"`
	Synonyms []string `json:"synonyms,omitempty"`
}

/*
{
 "description": "The object that contains individual type values.",
 "properties": {
  "synonyms": {
   "items": {
    "type": "string"
   },
   "type": "array"
  },
  "value": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
