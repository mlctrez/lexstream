package skill

import "time"

/*
SkillVersion Information about the skill version
*/
type SkillVersion struct {
	Version string `json,omitempty:"version"`
	/*
	   Description of the version (limited to 300 characters).
	*/
	Message      string    `json,omitempty:"message"`
	CreationTime time.Time `json,omitempty:"creationTime"`
	/*
	   List of submissions for the skill version
	*/
	Submissions []*VersionSubmission `json,omitempty:"submissions"`
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
