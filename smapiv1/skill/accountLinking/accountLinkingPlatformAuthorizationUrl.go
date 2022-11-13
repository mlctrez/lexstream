package accountlinking

/*
AccountLinkingPlatformAuthorizationUrl A key-value pair object that contains the OAuth2 authorization url to initiate the skill account linking process.
*/
type AccountLinkingPlatformAuthorizationUrl struct {
	PlatformType *PlatformType `json:"platformType"`
	// Defines the OAuth2 Authorization URL that will be used in this platform to authenticate the customer / person.
	PlatformAuthorizationUrl string `json:"platformAuthorizationUrl"`
}

/*
{
 "description": "A key-value pair object that contains the OAuth2 authorization url to initiate the skill account linking process.",
 "properties": {
  "platformAuthorizationUrl": {
   "description": "Defines the OAuth2 Authorization URL that will be used in this platform to authenticate the customer / person.",
   "type": "string"
  },
  "platformType": {
   "$ref": "#/definitions/v1.skill.accountLinking.PlatformType",
   "x-isEnum": true
  }
 },
 "required": [
  "platformAuthorizationUrl",
  "platformType"
 ],
 "type": "object"
}
*/
