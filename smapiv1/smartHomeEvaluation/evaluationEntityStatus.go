package smarthomeevaluation

type EvaluationEntityStatus string

func EvaluationEntityStatus_PASSED() EvaluationEntityStatus {
	return "PASSED"
}
func EvaluationEntityStatus_FAILED() EvaluationEntityStatus {
	return "FAILED"
}
func EvaluationEntityStatus_IN_PROGRESS() EvaluationEntityStatus {
	return "IN_PROGRESS"
}
func EvaluationEntityStatus_ERROR() EvaluationEntityStatus {
	return "ERROR"
}
