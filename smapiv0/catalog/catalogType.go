package catalog

/*
CatalogType Type of catalog.
*/
type CatalogType string

func CatalogType_AMAZONBroadcastChannel() CatalogType {
	return "AMAZON.BroadcastChannel"
}
func CatalogType_AMAZONGenre() CatalogType {
	return "AMAZON.Genre"
}
func CatalogType_AMAZONMusicAlbum() CatalogType {
	return "AMAZON.MusicAlbum"
}
func CatalogType_AMAZONMusicGroup() CatalogType {
	return "AMAZON.MusicGroup"
}
func CatalogType_AMAZONMusicPlaylist() CatalogType {
	return "AMAZON.MusicPlaylist"
}
func CatalogType_AMAZONMusicRecording() CatalogType {
	return "AMAZON.MusicRecording"
}
func CatalogType_AMAZONTerrestrialRadioChannel() CatalogType {
	return "AMAZON.TerrestrialRadioChannel"
}
func CatalogType_AMAZONAudioRecording() CatalogType {
	return "AMAZON.AudioRecording"
}
