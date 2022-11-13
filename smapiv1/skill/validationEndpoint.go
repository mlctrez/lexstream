package skill

/*
ValidationEndpoint Structure representing an endpoint.
*/
type ValidationEndpoint struct {
	// Path to the endpoint URI in the resource.
	PropertyPath string `json,omitempty:"propertyPath"`
	// Type of the endpoint (https, http, arn etc).
	Type_ string `json,omitempty:"type"`
	// Full URI of the endpoint.
	Value string `json,omitempty:"value"`
}

/*
{
 "description": "Structure representing an endpoint.",
 "properties": {
  "propertyPath": {
   "description": "Path to the endpoint URI in the resource.",
   "type": "string"
  },
  "type": {
   "description": "Type of the endpoint (https, http, arn etc).",
   "type": "string"
  },
  "value": {
   "description": "Full URI of the endpoint.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
