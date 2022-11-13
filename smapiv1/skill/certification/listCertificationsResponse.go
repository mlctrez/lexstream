package certification

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
ListCertificationsResponse List of certification summary for a skill.
*/
type ListCertificationsResponse struct {
	Links *smapiv1.Links `json,omitempty:"_links"`
	/*
	   boolean value for if the response is truncated. isTruncated = true if more than the assigned maxResults parameter value certification items are available for the skill. The results are then paginated and the remaining results can be retrieved in a similar paginated manner by using 'next' link in the _links or using the nextToken in a following request.
	*/
	IsTruncated bool `json,omitempty:"isTruncated"`
	// Encrypted token present when isTruncated is true.
	NextToken string `json,omitempty:"nextToken"`
	// Total number of certification results available for the skill.
	TotalCount int `json,omitempty:"totalCount"`
	/*
	   List of certifications available for a skill. The list of certifications is sorted in a default descending sort order on id field.
	*/
	Items []*CertificationSummary `json,omitempty:"items"`
}

/*
{
 "description": "List of certification summary for a skill.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "isTruncated": {
   "description": "boolean value for if the response is truncated. isTruncated = true if more than the assigned maxResults parameter value certification items are available for the skill. The results are then paginated and the remaining results can be retrieved in a similar paginated manner by using 'next' link in the _links or using the nextToken in a following request.\n",
   "type": "boolean"
  },
  "items": {
   "description": "List of certifications available for a skill. The list of certifications is sorted in a default descending sort order on id field.\n",
   "items": {
    "$ref": "#/definitions/v1.skill.certification.CertificationSummary"
   },
   "type": "array"
  },
  "nextToken": {
   "description": "Encrypted token present when isTruncated is true.",
   "type": "string"
  },
  "totalCount": {
   "description": "Total number of certification results available for the skill.",
   "type": "integer"
  }
 },
 "type": "object"
}
*/
