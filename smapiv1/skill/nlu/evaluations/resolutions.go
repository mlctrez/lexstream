package evaluations

/*
Resolutions A resolutions object representing the results of resolving the words captured from the user's utterance.
*/
type Resolutions struct {
	/*
	   An array of objects representing each possible authority for entity resolution. An authority represents the source for the data provided for the slot. For a custom slot type, the authority is the slot type you defined.
	*/
	ResolutionsPerAuthority map[string]ResolutionsPerAuthority `json:"resolutionsPerAuthority,omitempty"`
}

/*
{
 "description": "A resolutions object representing the results of resolving the words captured from the user's utterance.\n",
 "properties": {
  "resolutionsPerAuthority": {
   "description": "An array of objects representing each possible authority for entity resolution. An authority represents the source for the data provided for the slot. For a custom slot type, the authority is the slot type you defined.\n",
   "items": {
    "additionalProperties": {
     "$ref": "#/definitions/v1.skill.nlu.evaluations.ResolutionsPerAuthority"
    },
    "type": "object"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
