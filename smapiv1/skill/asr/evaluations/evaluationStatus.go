package evaluations

/*
EvaluationStatus Evaluation status:
  - `IN_PROGRESS` - indicate the evaluation is in progress.
  - `COMPLETED` - indicate the evaluation has been completed.
  - `FAILED` - indicate the evaluation has run into an error.
*/
type EvaluationStatus string

func EvaluationStatus_IN_PROGRESS() EvaluationStatus {
	return "IN_PROGRESS"
}
func EvaluationStatus_COMPLETED() EvaluationStatus {
	return "COMPLETED"
}
func EvaluationStatus_FAILED() EvaluationStatus {
	return "FAILED"
}
