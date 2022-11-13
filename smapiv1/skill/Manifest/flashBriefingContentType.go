package manifest

/*
FlashBriefingContentType format of the feed content.
*/
type FlashBriefingContentType string

func FlashBriefingContentType_TEXT() FlashBriefingContentType {
	return "TEXT"
}
func FlashBriefingContentType_AUDIO() FlashBriefingContentType {
	return "AUDIO"
}
func FlashBriefingContentType_AUDIO_AND_VIDEO() FlashBriefingContentType {
	return "AUDIO_AND_VIDEO"
}
