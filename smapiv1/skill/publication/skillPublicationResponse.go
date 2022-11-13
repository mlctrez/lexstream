package publication

import "time"

type SkillPublicationResponse struct {
	// Used to determine when the skill Publishing should start.
	PublishesAtDate time.Time               `json:"publishesAtDate"`
	Status          *SkillPublicationStatus `json:"status"`
}

/*
{
 "properties": {
  "publishesAtDate": {
   "description": "Used to determine when the skill Publishing should start.",
   "format": "date-time",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.publication.SkillPublicationStatus",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
