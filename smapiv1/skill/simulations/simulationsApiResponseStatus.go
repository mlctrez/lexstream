package simulations

/*
SimulationsApiResponseStatus String that specifies the current status of the simulation. Possible values are "IN_PROGRESS", "SUCCESSFUL", and "FAILED".
*/
type SimulationsApiResponseStatus string

func SimulationsApiResponseStatus_IN_PROGRESS() SimulationsApiResponseStatus {
	return "IN_PROGRESS"
}
func SimulationsApiResponseStatus_SUCCESSFUL() SimulationsApiResponseStatus {
	return "SUCCESSFUL"
}
func SimulationsApiResponseStatus_FAILED() SimulationsApiResponseStatus {
	return "FAILED"
}
