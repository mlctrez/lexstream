package interactionmodel

/*
FallbackIntentSensitivityLevel Skill's sensitivity level for out-of-domain utterances. By default, the sensitivity level of the skill is set to ‘LOW’.
As the sensitivity level for a skill is increased, more customer utterances that are not supported by the skill will be captured by AMAZON.FallbackIntent.
*/
type FallbackIntentSensitivityLevel string

func FallbackIntentSensitivityLevel_HIGH() FallbackIntentSensitivityLevel {
	return "HIGH"
}
func FallbackIntentSensitivityLevel_MEDIUM() FallbackIntentSensitivityLevel {
	return "MEDIUM"
}
func FallbackIntentSensitivityLevel_LOW() FallbackIntentSensitivityLevel {
	return "LOW"
}
