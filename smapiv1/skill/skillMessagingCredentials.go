package skill

/*
SkillMessagingCredentials Defines the structure for skill messaging credentials.
*/
type SkillMessagingCredentials struct {
	// The client id for the security profile.
	ClientId string `json,omitempty:"clientId"`
	// The client secret for the security profile.
	ClientSecret string `json,omitempty:"clientSecret"`
}

/*
{
 "description": "Defines the structure for skill messaging credentials.",
 "properties": {
  "clientId": {
   "description": "The client id for the security profile.",
   "type": "string"
  },
  "clientSecret": {
   "description": "The client secret for the security profile.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
