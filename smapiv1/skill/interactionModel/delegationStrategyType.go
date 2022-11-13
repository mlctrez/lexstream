package interactionmodel

/*
DelegationStrategyType Enumerates delegation strategies used to control automatic dialog management through the defined dialog model. When no delegation strategies are defined, the value SKILL_RESPONSE is assumed.
*/
type DelegationStrategyType string

func DelegationStrategyType_ALWAYS() DelegationStrategyType {
	return "ALWAYS"
}
func DelegationStrategyType_SKILL_RESPONSE() DelegationStrategyType {
	return "SKILL_RESPONSE"
}
