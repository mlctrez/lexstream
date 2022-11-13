package private

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
ListPrivateDistributionAccountsResponse Response of ListPrivateDistributionAccounts.
*/
type ListPrivateDistributionAccountsResponse struct {
	Links *smapiv1.Links `json:"_links,omitempty"`
	// List of PrivateDistributionAccounts.
	PrivateDistributionAccounts []*PrivateDistributionAccount `json:"privateDistributionAccounts,omitempty"`
	NextToken                   string                        `json:"nextToken,omitempty"`
}

/*
{
 "description": "Response of ListPrivateDistributionAccounts.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "nextToken": {
   "type": "string"
  },
  "privateDistributionAccounts": {
   "description": "List of PrivateDistributionAccounts.",
   "items": {
    "$ref": "#/definitions/v1.skill.Private.PrivateDistributionAccount"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/
