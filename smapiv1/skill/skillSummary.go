package skill

import (
	smapiv1 "github.com/mlctrez/lexstream/smapiv1"
	"time"
)

/*
SkillSummary Information about the skills.
*/
type SkillSummary struct {
	SkillId string               `json,omitempty:"skillId"`
	Stage   *smapiv1.StageV2Type `json,omitempty:"stage"`
	// List of APIs currently implemented by the skill.
	Apis              []*SkillSummaryApis `json,omitempty:"apis"`
	PublicationStatus *PublicationStatus  `json,omitempty:"publicationStatus"`
	LastUpdated       time.Time           `json,omitempty:"lastUpdated"`
	/*
	   Name of the skill in skill locales (keys are locale names (e.g. 'en-US') whereas values are name of the
	   skill in that locale.
	*/
	NameByLocale map[string]string `json,omitempty:"nameByLocale"`
	/*
	   Amazon Standard Identification Number (ASIN) is unique blocks of 10 letters and/or numbers that identify items. More info about ASIN can be found here: https://www.amazon.com/gp/seller/asin-upc-isbn-info.html
	   ASIN is available for those skills only, that have been published, at least once.
	*/
	Asin  string         `json,omitempty:"asin"`
	Links *smapiv1.Links `json,omitempty:"_links"`
}

/*
{
 "description": "Information about the skills.",
 "properties": {
  "_links": {
   "$ref": "#/definitions/v1.Links"
  },
  "apis": {
   "description": "List of APIs currently implemented by the skill.",
   "items": {
    "$ref": "#/definitions/v1.skill.SkillSummaryApis"
   },
   "type": "array"
  },
  "asin": {
   "description": "Amazon Standard Identification Number (ASIN) is unique blocks of 10 letters and/or numbers that identify items. More info about ASIN can be found here: https://www.amazon.com/gp/seller/asin-upc-isbn-info.html\nASIN is available for those skills only, that have been published, at least once.\n",
   "type": "string"
  },
  "lastUpdated": {
   "format": "date-time",
   "type": "string"
  },
  "nameByLocale": {
   "additionalProperties": {
    "type": "string"
   },
   "description": "Name of the skill in skill locales (keys are locale names (e.g. 'en-US') whereas values are name of the\nskill in that locale.\n",
   "type": "object"
  },
  "publicationStatus": {
   "$ref": "#/definitions/v1.skill.PublicationStatus",
   "x-isEnum": true
  },
  "skillId": {
   "type": "string"
  },
  "stage": {
   "$ref": "#/definitions/v1.StageV2Type",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
