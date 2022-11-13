package certification

/*
PublicationFailure Information about why the skill was not published in certain countries.
*/
type PublicationFailure struct {
	// Reason why Amazon did not publish the skill in certain countries.
	Reason string `json:"reason"`
	// List of countries where Amazon did not publish the skill for a specific reason
	Countries []string `json:"countries"`
}

/*
{
 "description": "Information about why the skill was not published in certain countries.",
 "properties": {
  "countries": {
   "description": "List of countries where Amazon did not publish the skill for a specific reason",
   "items": {
    "description": "Two letter country codes in ISO 3166-1 alpha-2 format.",
    "type": "string"
   },
   "type": "array"
  },
  "reason": {
   "description": "Reason why Amazon did not publish the skill in certain countries.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
