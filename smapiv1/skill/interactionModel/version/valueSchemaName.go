package version

type ValueSchemaName struct {
	Value    string   `json:"value,omitempty"`
	Synonyms []string `json:"synonyms,omitempty"`
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
