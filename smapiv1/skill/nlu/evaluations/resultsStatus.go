package evaluations

type ResultsStatus string

func ResultsStatus_PASSED() ResultsStatus {
	return "PASSED"
}
func ResultsStatus_FAILED() ResultsStatus {
	return "FAILED"
}
