package certification

/*
DistributionInfo The distribution information for skill where Amazon distributed the skill
*/
type DistributionInfo struct {
	// All countries where the skill was published in by Amazon.
	PublishedCountries  []string              `json,omitempty:"publishedCountries"`
	PublicationFailures []*PublicationFailure `json,omitempty:"publicationFailures"`
}

/*
{
 "description": "The distribution information for skill where Amazon distributed the skill",
 "properties": {
  "publicationFailures": {
   "items": {
    "$ref": "#/definitions/v1.skill.certification.PublicationFailure"
   },
   "type": "array"
  },
  "publishedCountries": {
   "description": "All countries where the skill was published in by Amazon.",
   "items": {
    "description": "Two letter country codes in ISO 3166-1 alpha-2 format.",
    "type": "string"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
