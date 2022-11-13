package manifest

/*
MusicContentName Name of the content type that's supported for the music skill.
*/
type MusicContentName string

func MusicContentName_ON_DEMAND() MusicContentName {
	return "ON_DEMAND"
}
func MusicContentName_RADIO() MusicContentName {
	return "RADIO"
}
func MusicContentName_PODCAST() MusicContentName {
	return "PODCAST"
}
