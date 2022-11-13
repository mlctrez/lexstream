package skill

/*
HostedSkillDeploymentDetails Details about hosted skill deployment.
*/
type HostedSkillDeploymentDetails struct {
	CommitId string `json:"commitId,omitempty"`
	LogUrl   string `json:"logUrl,omitempty"`
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
