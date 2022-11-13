package version

type ValueSchemaName struct {
	Value    string   `json:"value"`
	Synonyms []string `json:"synonyms"`
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
