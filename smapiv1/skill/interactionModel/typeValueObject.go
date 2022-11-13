package interactionmodel

/*
TypeValueObject The object that contains individual type values.
*/
type TypeValueObject struct {
	Value    string   `json,omitempty:"value"`
	Synonyms []string `json,omitempty:"synonyms"`
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
