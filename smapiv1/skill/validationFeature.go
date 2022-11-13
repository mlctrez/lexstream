package skill

/*
ValidationFeature Structure representing a public feature.
*/
type ValidationFeature struct {
	// Name of the feature.
	Name string `json:"name,omitempty"`
	// Contact URL or email for the feature.
	Contact string `json:"contact,omitempty"`
}

/*
{
 "description": "Structure representing a public feature.",
 "properties": {
  "contact": {
   "description": "Contact URL or email for the feature.",
   "type": "string"
  },
  "name": {
   "description": "Name of the feature.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
