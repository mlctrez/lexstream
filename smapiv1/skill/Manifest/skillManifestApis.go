package manifest

/*
SkillManifestApis Defines the structure for implemented apis information in the skill manifest.
*/
type SkillManifestApis struct {
	FlashBriefing    *FlashBriefingApis    `json,omitempty:"flashBriefing"`
	Custom           *CustomApis           `json,omitempty:"custom"`
	SmartHome        *SmartHomeApis        `json,omitempty:"smartHome"`
	Video            *VideoApis            `json,omitempty:"video"`
	AlexaForBusiness *AlexaForBusinessApis `json,omitempty:"alexaForBusiness"`
	Health           *HealthApis           `json,omitempty:"health"`
	HouseholdList    *HouseHoldList        `json,omitempty:"householdList"`
	Music            *MusicApis            `json,omitempty:"music"`
}

/*
{
 "description": "Defines the structure for implemented apis information in the skill manifest.",
 "properties": {
  "alexaForBusiness": {
   "$ref": "#/definitions/v1.skill.Manifest.AlexaForBusinessApis"
  },
  "custom": {
   "$ref": "#/definitions/v1.skill.Manifest.CustomApis"
  },
  "flashBriefing": {
   "$ref": "#/definitions/v1.skill.Manifest.FlashBriefingApis"
  },
  "health": {
   "$ref": "#/definitions/v1.skill.Manifest.HealthApis"
  },
  "householdList": {
   "$ref": "#/definitions/v1.skill.Manifest.HouseHoldList"
  },
  "music": {
   "$ref": "#/definitions/v1.skill.Manifest.MusicApis"
  },
  "smartHome": {
   "$ref": "#/definitions/v1.skill.Manifest.SmartHomeApis"
  },
  "video": {
   "$ref": "#/definitions/v1.skill.Manifest.VideoApis"
  }
 },
 "type": "object"
}
*/
