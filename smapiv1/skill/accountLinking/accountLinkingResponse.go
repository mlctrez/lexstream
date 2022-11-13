package accountlinking

/*
AccountLinkingResponse The account linking information of a skill.
*/
type AccountLinkingResponse struct {
	Type_ *AccountLinkingType `json:"type,omitempty"`
	// The url where customers will be redirected in the companion app to enter login credentials.
	AuthorizationUrl string `json:"authorizationUrl,omitempty"`
	// The list of domains that the authorization URL will fetch content from.
	Domains []string `json:"domains,omitempty"`
	// The unique public string used to identify the client requesting for authentication.
	ClientId string `json:"clientId,omitempty"`
	// The list of permissions which will be requested from the skill user.
	Scopes []string `json:"scopes,omitempty"`
	// The url used for access token and token refresh requests.
	AccessTokenUrl    string                 `json:"accessTokenUrl,omitempty"`
	AccessTokenScheme *AccessTokenSchemeType `json:"accessTokenScheme,omitempty"`
	/*
	   The time in seconds for which access token is valid.
	   If OAuth client returns "expires_in", it will be overwrite this parameter.
	*/
	DefaultTokenExpirationInSeconds int `json:"defaultTokenExpirationInSeconds,omitempty"`
	// The list of valid urls to redirect back to, when the linking process is initiated from a third party system.
	RedirectUrls []string `json:"redirectUrls,omitempty"`
	// The list of valid authorization urls for allowed platforms to initiate account linking.
	AuthorizationUrlsByPlatform []*AccountLinkingPlatformAuthorizationUrl `json:"authorizationUrlsByPlatform,omitempty"`
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
