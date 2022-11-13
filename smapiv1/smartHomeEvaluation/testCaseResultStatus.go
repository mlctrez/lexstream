package smarthomeevaluation

type TestCaseResultStatus string

func TestCaseResultStatus_PASSED() TestCaseResultStatus {
	return "PASSED"
}
func TestCaseResultStatus_FAILED() TestCaseResultStatus {
	return "FAILED"
}
func TestCaseResultStatus_ERROR() TestCaseResultStatus {
	return "ERROR"
}
