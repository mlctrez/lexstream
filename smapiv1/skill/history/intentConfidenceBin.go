package history

/*
IntentConfidenceBin A filter used to retrieve items where the intent confidence bin is equal to the given value.
* `HIGH`: Intent was recognized with high confidence.
* `MEDIUM`: Intent was recognized with medium confidence.
* `LOW`: Intent was recognized with low confidence. Note: Low confidence intents are not sent to the skill.
*/
type IntentConfidenceBin string

func IntentConfidenceBin_HIGH() IntentConfidenceBin {
	return "HIGH"
}
func IntentConfidenceBin_MEDIUM() IntentConfidenceBin {
	return "MEDIUM"
}
func IntentConfidenceBin_LOW() IntentConfidenceBin {
	return "LOW"
}
