package skill

import "time"

/*
SkillVersion Information about the skill version
*/
type SkillVersion struct {
	Version string `json:"version,omitempty"`
	/*
	   Description of the version (limited to 300 characters).
	*/
	Message      string    `json:"message,omitempty"`
	CreationTime time.Time `json:"creationTime,omitempty"`
	/*
	   List of submissions for the skill version
	*/
	Submissions []*VersionSubmission `json:"submissions,omitempty"`
}

/*
{
 "description": "Information about the skill version",
 "properties": {
  "creationTime": {
   "format": "date-time",
   "type": "string"
  },
  "message": {
   "description": "Description of the version (limited to 300 characters).\n",
   "type": "string"
  },
  "submissions": {
   "description": "List of submissions for the skill version\n",
   "items": {
    "$ref": "#/definitions/v1.skill.VersionSubmission"
   },
   "type": "array"
  },
  "version": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
