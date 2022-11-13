package invocations

/*
InvocationResponseStatus String that specifies the status of skill invocation. Possible values are "SUCCEEDED", and "FAILED".
*/
type InvocationResponseStatus string

func InvocationResponseStatus_SUCCEEDED() InvocationResponseStatus {
	return "SUCCEEDED"
}
func InvocationResponseStatus_FAILED() InvocationResponseStatus {
	return "FAILED"
}
