package manifest

/*
DistributionMode What audience the skill should be distributed to. "PUBLIC" - available to all users. Has ASIN and can be enabled. "PRIVATE" - available to entitled users. Has ASIN and can be enabled. "INTERNAL" - has no ASIN and cannot be enabled by users. Internally managed skills.
*/
type DistributionMode string

func DistributionMode_PRIVATE() DistributionMode {
	return "PRIVATE"
}
func DistributionMode_PUBLIC() DistributionMode {
	return "PUBLIC"
}
func DistributionMode_INTERNAL() DistributionMode {
	return "INTERNAL"
}
