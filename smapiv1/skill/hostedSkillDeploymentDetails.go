package skill

/*
HostedSkillDeploymentDetails Details about hosted skill deployment.
*/
type HostedSkillDeploymentDetails struct {
	CommitId string `json:"commitId"`
	LogUrl   string `json:"logUrl"`
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
