package smapiv1

type StageV2Type string

func StageV2Type_Live() StageV2Type {
	return "live"
}
func StageV2Type_Certified() StageV2Type {
	return "certified"
}
func StageV2Type_Development() StageV2Type {
	return "development"
}
