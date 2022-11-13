package smarthomeevaluation

type SHCapabilityErrorCode string

func SHCapabilityErrorCode_NO_SUCH_ENDPOINT() SHCapabilityErrorCode {
	return "NO_SUCH_ENDPOINT"
}
func SHCapabilityErrorCode_NO_SUCH_SKILL_STAGE() SHCapabilityErrorCode {
	return "NO_SUCH_SKILL_STAGE"
}
func SHCapabilityErrorCode_NO_SUCH_TEST_PLAN() SHCapabilityErrorCode {
	return "NO_SUCH_TEST_PLAN"
}
func SHCapabilityErrorCode_MULTIPLE_MATCHED_ENDPOINTS() SHCapabilityErrorCode {
	return "MULTIPLE_MATCHED_ENDPOINTS"
}
func SHCapabilityErrorCode_MULTIPLE_MATCHED_TEST_PLANS() SHCapabilityErrorCode {
	return "MULTIPLE_MATCHED_TEST_PLANS"
}
func SHCapabilityErrorCode_CAPABILITY_NOT_SUPPORTED() SHCapabilityErrorCode {
	return "CAPABILITY_NOT_SUPPORTED"
}
func SHCapabilityErrorCode_DISCOVERY_FAILED() SHCapabilityErrorCode {
	return "DISCOVERY_FAILED"
}
func SHCapabilityErrorCode_TEST_CASE_TIME_OUT() SHCapabilityErrorCode {
	return "TEST_CASE_TIME_OUT"
}
