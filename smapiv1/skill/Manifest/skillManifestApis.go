package manifest

/*
SkillManifestApis Defines the structure for implemented apis information in the skill manifest.
*/
type SkillManifestApis struct {
	FlashBriefing    *FlashBriefingApis    `json:"flashBriefing"`
	Custom           *CustomApis           `json:"custom"`
	SmartHome        *SmartHomeApis        `json:"smartHome"`
	Video            *VideoApis            `json:"video"`
	AlexaForBusiness *AlexaForBusinessApis `json:"alexaForBusiness"`
	Health           *HealthApis           `json:"health"`
	HouseholdList    *HouseHoldList        `json:"householdList"`
	Music            *MusicApis            `json:"music"`
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
