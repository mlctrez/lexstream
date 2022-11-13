package isp

/*
Stage Stage of in-skill product.
*/
type Stage string

func Stage_Development() Stage {
	return "development"
}
func Stage_Live() Stage {
	return "live"
}
