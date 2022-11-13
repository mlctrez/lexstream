package accountlinking

/*
AccountLinkingRequestPayload The payload for creating the account linking partner.
*/
type AccountLinkingRequestPayload struct {
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
	AccessTokenUrl string `json,omitempty:"accessTokenUrl"`
	// The client secret provided by developer.
	ClientSecret      string                 `json,omitempty:"clientSecret"`
	AccessTokenScheme *AccessTokenSchemeType `json,omitempty:"accessTokenScheme"`
	/*
	   The time in seconds for which access token is valid.
	   If OAuth client returns "expires_in", it will be overwrite this parameter.
	*/
	DefaultTokenExpirationInSeconds int `json,omitempty:"defaultTokenExpirationInSeconds"`
	// Optional, if your skill requires reciprocal authorization, provide this additional access token url to handle reciprocal (Alexa) authorization codes that can be exchanged for Alexa access tokens.
	ReciprocalAccessTokenUrl string `json,omitempty:"reciprocalAccessTokenUrl"`
	// The list of valid urls to redirect back to, when the linking process is initiated from a third party system.
	RedirectUrls []string `json,omitempty:"redirectUrls"`
	// The list of valid authorization urls for allowed platforms to initiate account linking.
	AuthorizationUrlsByPlatform []*AccountLinkingPlatformAuthorizationUrl `json,omitempty:"authorizationUrlsByPlatform"`
	// Set to true to let users enable the skill without starting the account linking flow. Set to false to require the normal account linking flow when users enable the skill.
	SkipOnEnablement bool `json,omitempty:"skipOnEnablement"`
}

/*
{
 "description": "The payload for creating the account linking partner.",
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
  "clientSecret": {
   "description": "The client secret provided by developer.",
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
  "reciprocalAccessTokenUrl": {
   "description": "Optional, if your skill requires reciprocal authorization, provide this additional access token url to handle reciprocal (Alexa) authorization codes that can be exchanged for Alexa access tokens.",
   "type": "string"
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
  "skipOnEnablement": {
   "description": "Set to true to let users enable the skill without starting the account linking flow. Set to false to require the normal account linking flow when users enable the skill.",
   "type": "boolean"
  },
  "type": {
   "$ref": "#/definitions/v1.skill.accountLinking.AccountLinkingType",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
