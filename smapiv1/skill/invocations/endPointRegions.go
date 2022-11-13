package invocations

/*
EndPointRegions Region of endpoint to be called.
*/
type EndPointRegions string

func EndPointRegions_NA() EndPointRegions {
	return "NA"
}
func EndPointRegions_EU() EndPointRegions {
	return "EU"
}
func EndPointRegions_FE() EndPointRegions {
	return "FE"
}
