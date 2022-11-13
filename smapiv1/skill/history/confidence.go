package history

/*
Confidence The hypothesized confidence for this interaction.
*/
type Confidence struct {
	Bin *ConfidenceBin `json:"bin"`
}

/*
{
 "description": "The hypothesized confidence for this interaction.",
 "properties": {
  "bin": {
   "$ref": "#/definitions/v1.skill.history.ConfidenceBin",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
