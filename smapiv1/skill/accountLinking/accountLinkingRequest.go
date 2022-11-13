package accountlinking

/*
AccountLinkingRequest The request body of AccountLinkingRequest.
*/
type AccountLinkingRequest struct {
	AccountLinkingRequest *AccountLinkingRequestPayload `json:"accountLinkingRequest"`
}

/*
{
 "description": "The request body of AccountLinkingRequest.",
 "properties": {
  "accountLinkingRequest": {
   "$ref": "#/definitions/v1.skill.accountLinking.AccountLinkingRequestPayload"
  }
 },
 "type": "object"
}
*/
