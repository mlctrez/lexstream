package alexahosted

import "time"

type HostedSkillRepositoryCredentials struct {
	// AWS Access Key ID used to access hosted skill repository
	Username string `json,omitempty:"username"`
	// signed AWS Credentials used to access hosted skill repository
	Password string `json,omitempty:"password"`
	// expiration time for the credentials
	ExpiresAt time.Time `json,omitempty:"expiresAt"`
}

/*
{
 "properties": {
  "expiresAt": {
   "description": "expiration time for the credentials",
   "format": "date-time",
   "type": "string"
  },
  "password": {
   "description": "signed AWS Credentials used to access hosted skill repository",
   "type": "string"
  },
  "username": {
   "description": "AWS Access Key ID used to access hosted skill repository",
   "type": "string"
  }
 },
 "type": "object"
}
*/
