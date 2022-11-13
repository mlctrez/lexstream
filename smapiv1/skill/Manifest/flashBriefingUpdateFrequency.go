package manifest

/*
FlashBriefingUpdateFrequency Tells how often the feed has new content.
*/
type FlashBriefingUpdateFrequency string

func FlashBriefingUpdateFrequency_HOURLY() FlashBriefingUpdateFrequency {
	return "HOURLY"
}
func FlashBriefingUpdateFrequency_DAILY() FlashBriefingUpdateFrequency {
	return "DAILY"
}
func FlashBriefingUpdateFrequency_WEEKLY() FlashBriefingUpdateFrequency {
	return "WEEKLY"
}
func FlashBriefingUpdateFrequency_UNKNOWN() FlashBriefingUpdateFrequency {
	return "UNKNOWN"
}
