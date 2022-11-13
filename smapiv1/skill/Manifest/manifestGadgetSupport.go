package manifest

/*
ManifestGadgetSupport Defines the structure for gadget buttons support in the skill manifest.
*/
type ManifestGadgetSupport struct {
	Requirement *GadgetSupport `json:"requirement,omitempty"`
	// Minimum number of gadget buttons required.
	MinGadgetButtons int `json:"minGadgetButtons,omitempty"`
	// Maximum number of gadget buttons required.
	MaxGadgetButtons int `json:"maxGadgetButtons,omitempty"`
	// Maximum number of players in the game.
	NumPlayersMax int `json:"numPlayersMax,omitempty"`
	// Minimum number of players in the game.
	NumPlayersMin int `json:"numPlayersMin,omitempty"`
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
