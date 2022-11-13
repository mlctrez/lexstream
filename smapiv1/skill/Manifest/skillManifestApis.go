package manifest

/*
SkillManifestApis Defines the structure for implemented apis information in the skill manifest.
*/
type SkillManifestApis struct {
	FlashBriefing    *FlashBriefingApis    `json:"flashBriefing,omitempty"`
	Custom           *CustomApis           `json:"custom,omitempty"`
	SmartHome        *SmartHomeApis        `json:"smartHome,omitempty"`
	Video            *VideoApis            `json:"video,omitempty"`
	AlexaForBusiness *AlexaForBusinessApis `json:"alexaForBusiness,omitempty"`
	Health           *HealthApis           `json:"health,omitempty"`
	HouseholdList    *HouseHoldList        `json:"householdList,omitempty"`
	Music            *MusicApis            `json:"music,omitempty"`
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
