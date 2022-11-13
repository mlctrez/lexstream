package history

/*
ConfidenceBin Intent confidence bin for this utterance.
* `HIGH`: Intent was recognized with high confidence.
* `MEDIUM`: Intent was recognized with medium confidence.
* `LOW`: Intent was recognized with low confidence. Note: Low confidence intents are not sent to the skill.
*/
type ConfidenceBin string

func ConfidenceBin_HIGH() ConfidenceBin {
	return "HIGH"
}
func ConfidenceBin_MEDIUM() ConfidenceBin {
	return "MEDIUM"
}
func ConfidenceBin_LOW() ConfidenceBin {
	return "LOW"
}
