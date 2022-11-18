package amsapi

import "github.com/mlctrez/lexstream/amsapi/types"

// MetadataName is used for voice prompt or display use cases of entity (artist, song, etc.) names.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#metadatanameproperty
type MetadataName struct {
	Speech  SpeechInfo `json:"speech"`
	Display string     `json:"display"`
}

// BuildMetadataName is shorthand for:
//
//	return MetadataName{
//		Speech: SpeechInfo{
//			Type: SpeechInfoTypePlainText,
//			Text: text,
//		},
//		Display: display,
//	}
func BuildMetadataName(text, display string) MetadataName {
	return MetadataName{
		Speech: SpeechInfo{
			Type: SpeechInfoTypePlainText,
			Text: text,
		},
		Display: display,
	}
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
	Type    types.EntityType `json:"type"`
	Name    MetadataName     `json:"name"`
	Art     Art              `json:"art"`
	Authors []EntityMetadata `json:"authors,omitempty"`
	Album   EntityMetadata   `json:"albums,omitempty"`
	Artists []EntityMetadata `json:"artists,omitempty"`
}

// EntityMetadata is for entities like artists, songs, etc.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#entitymetadata
type EntityMetadata struct {
	Name MetadataNameProperty `json:"name"`
}

// BuildAuthors builds a slice of one author with the given text and display data.
func BuildAuthors(text, display string) []EntityMetadata {
	return []EntityMetadata{{
		Name: MetadataNameProperty{
			Speech: SpeechInfo{
				Type: SpeechInfoTypePlainText,
				Text: text},
			Display: display,
		},
	}}
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
	Type SpeechInfoType `json:"type"`
	Text string         `json:"text"`
}

type SpeechInfoType string

const SpeechInfoTypePlainText SpeechInfoType = "PLAIN_TEXT"
