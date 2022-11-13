package version

type ValueSchemaName struct {
	Value    string   `json,omitempty:"value"`
	Synonyms []string `json,omitempty:"synonyms"`
}

/*
{
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
