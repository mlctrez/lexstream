package accountlinking

/*
AccountLinkingResponse The account linking information of a skill.
*/
type AccountLinkingResponse struct {
	Type_ *AccountLinkingType `json,omitempty:"type"`
	// The url where customers will be redirected in the companion app to enter login credentials.
	AuthorizationUrl string `json,omitempty:"authorizationUrl"`
	// The list of domains that the authorization URL will fetch content from.
	Domains []string `json,omitempty:"domains"`
	// The unique public string used to identify the client requesting for authentication.
	ClientId string `json,omitempty:"clientId"`
	// The list of permissions which will be requested from the skill user.
	Scopes []string `json,omitempty:"scopes"`
	// The url used for access token and token refresh requests.
	AccessTokenUrl    string                 `json,omitempty:"accessTokenUrl"`
	AccessTokenScheme *AccessTokenSchemeType `json,omitempty:"accessTokenScheme"`
	/*
	   The time in seconds for which access token is valid.
	   If OAuth client returns "expires_in", it will be overwrite this parameter.
	*/
	DefaultTokenExpirationInSeconds int `json,omitempty:"defaultTokenExpirationInSeconds"`
	// The list of valid urls to redirect back to, when the linking process is initiated from a third party system.
	RedirectUrls []string `json,omitempty:"redirectUrls"`
	// The list of valid authorization urls for allowed platforms to initiate account linking.
	AuthorizationUrlsByPlatform []*AccountLinkingPlatformAuthorizationUrl `json,omitempty:"authorizationUrlsByPlatform"`
}

/*
{
 "description": "The account linking information of a skill.",
 "properties": {
  "accessTokenScheme": {
   "$ref": "#/definitions/v1.skill.accountLinking.AccessTokenSchemeType",
   "x-isEnum": true
  },
  "accessTokenUrl": {
   "description": "The url used for access token and token refresh requests.",
   "type": "string"
  },
  "authorizationUrl": {
   "description": "The url where customers will be redirected in the companion app to enter login credentials.",
   "format": "uri",
   "type": "string"
  },
  "authorizationUrlsByPlatform": {
   "description": "The list of valid authorization urls for allowed platforms to initiate account linking.",
   "items": {
    "$ref": "#/definitions/v1.skill.accountLinking.AccountLinkingPlatformAuthorizationUrl"
   },
   "type": "array"
  },
  "clientId": {
   "description": "The unique public string used to identify the client requesting for authentication.",
   "type": "string"
  },
  "defaultTokenExpirationInSeconds": {
   "description": "The time in seconds for which access token is valid.\nIf OAuth client returns \"expires_in\", it will be overwrite this parameter.\n",
   "type": "integer"
  },
  "domains": {
   "description": "The list of domains that the authorization URL will fetch content from.",
   "items": {
    "type": "string"
   },
   "type": "array"
  },
  "redirectUrls": {
   "description": "The list of valid urls to redirect back to, when the linking process is initiated from a third party system.",
   "items": {
    "type": "string"
   },
   "type": "array"
  },
  "scopes": {
   "description": "The list of permissions which will be requested from the skill user.",
   "items": {
    "type": "string"
   },
   "type": "array"
  },
  "type": {
   "$ref": "#/definitions/v1.skill.accountLinking.AccountLinkingType",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
