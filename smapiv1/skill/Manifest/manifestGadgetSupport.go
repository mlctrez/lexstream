package manifest

/*
ManifestGadgetSupport Defines the structure for gadget buttons support in the skill manifest.
*/
type ManifestGadgetSupport struct {
	Requirement *GadgetSupport `json:"requirement"`
	// Minimum number of gadget buttons required.
	MinGadgetButtons int `json:"minGadgetButtons"`
	// Maximum number of gadget buttons required.
	MaxGadgetButtons int `json:"maxGadgetButtons"`
	// Maximum number of players in the game.
	NumPlayersMax int `json:"numPlayersMax"`
	// Minimum number of players in the game.
	NumPlayersMin int `json:"numPlayersMin"`
}

/*
{
 "description": "Defines the structure for gadget buttons support in the skill manifest.",
 "properties": {
  "maxGadgetButtons": {
   "description": "Maximum number of gadget buttons required.",
   "maximum": 4,
   "minimum": 1,
   "type": "integer"
  },
  "minGadgetButtons": {
   "description": "Minimum number of gadget buttons required.",
   "maximum": 4,
   "minimum": 1,
   "type": "integer"
  },
  "numPlayersMax": {
   "description": "Maximum number of players in the game.",
   "maximum": 16,
   "minimum": 1,
   "type": "integer"
  },
  "numPlayersMin": {
   "description": "Minimum number of players in the game.",
   "maximum": 16,
   "minimum": 1,
   "type": "integer"
  },
  "requirement": {
   "$ref": "#/definitions/v1.skill.Manifest.GadgetSupport",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
