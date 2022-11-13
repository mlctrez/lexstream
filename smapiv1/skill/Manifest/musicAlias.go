package manifest

type MusicAlias struct {
	// Alias name to be associated with the music skill.
	Name string `json:"name,omitempty"`
}

/*
{
 "properties": {
  "name": {
   "description": "Alias name to be associated with the music skill.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
