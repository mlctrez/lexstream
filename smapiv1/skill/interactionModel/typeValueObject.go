package interactionmodel

/*
TypeValueObject The object that contains individual type values.
*/
type TypeValueObject struct {
	Value    string   `json:"value"`
	Synonyms []string `json:"synonyms"`
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
