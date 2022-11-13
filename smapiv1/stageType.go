package smapiv1

type StageType string

func StageType_Development() StageType {
	return "development"
}
func StageType_Live() StageType {
	return "live"
}
