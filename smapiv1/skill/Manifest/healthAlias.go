package manifest

type HealthAlias struct {
	// Name of alias to use when invoking a health skill.
	Name string `json:"name"`
}

/*
{
 "properties": {
  "name": {
   "description": "Name of alias to use when invoking a health skill.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
