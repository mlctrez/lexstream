package evaluations

/*
EvaluationResultStatus enum indicating the evaluation result status.
  - `PASSED` - evaluation result is considered passed
  - `FAILED` - evaluation result is considered failed
*/
type EvaluationResultStatus string

func EvaluationResultStatus_PASSED() EvaluationResultStatus {
	return "PASSED"
}
func EvaluationResultStatus_FAILED() EvaluationResultStatus {
	return "FAILED"
}
