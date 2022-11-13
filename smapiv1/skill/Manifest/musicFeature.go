package manifest

type MusicFeature struct {
	// Feature name to be associated with the music skill.
	Name string `json,omitempty:"name"`
}

/*
{
 "properties": {
  "name": {
   "description": "Feature name to be associated with the music skill.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
