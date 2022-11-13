package skill

/*
HostedSkillDeploymentDetails Details about hosted skill deployment.
*/
type HostedSkillDeploymentDetails struct {
	CommitId string `json,omitempty:"commitId"`
	LogUrl   string `json,omitempty:"logUrl"`
}

/*
{
 "description": "Details about hosted skill deployment.",
 "properties": {
  "commitId": {
   "type": "string"
  },
  "logUrl": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
