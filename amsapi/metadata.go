package amsapi

// MetadataName is used for voice prompt or display use cases of entity (artist, song, etc.) names.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#metadatanameproperty
type MetadataName struct {
	Speech  SpeechInfo `json:"speech"`
	Display string     `json:"display"`
}

// MediaMetadata contains metadata (for example, album name, author name, title, "type", etc.) for a media
// item (i.e., Content or Item). MediaMetadata extends BaseMetadata, and supports the following metadata types:
//
//	ALBUM - for an album (Music only)
//	ARTIST - for an artist (Music only)
//	GENRE - for a genre (Music only)
//	PLAYLIST - for a playlist
//	PROGRAM - for a program (episode) or program series (Podcast only)
//	STATION - for a station (Radio only)
//	TRACK - for a track/song (Music only)
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#mediametadata
type MediaMetadata struct {
	Type    string           `json:"type"`
	Name    MetadataName     `json:"name"`
	Art     Art              `json:"art"`
	Authors []EntityMetadata `json:"authors"`
	Album   EntityMetadata   `json:"albums,omitempty"`
	Artists []EntityMetadata `json:"artists,omitempty"`
}

// EntityMetadata is for entities like artists, songs, etc.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#entitymetadata
type EntityMetadata struct {
	Name MetadataNameProperty `json:"name"`
}

// MetadataNameProperty is used for voice prompt or display use cases of entity (artist, song, etc.) names.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#metadatanameproperty
type MetadataNameProperty struct {
	Speech  SpeechInfo `json:"speech"`
	Display string     `json:"display"`
}

// SpeechInfo captures the details of how Alexa should render text in voice prompts to the user.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#speechinfo
type SpeechInfo struct {
	Type string `json:"type"`
	Text string `json:"text"`
}