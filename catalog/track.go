package catalog

import (
	"fmt"
	"time"
)

func CreateTrackCatalog() *TrackCatalog {
	return &TrackCatalog{Header: buildHeader(MusicRecording)}
}

type TrackCatalog struct {
	Header
	Entities []TrackEntity `json:"entities"`
}

type TrackEntity struct {
	Id                string          `json:"id"`
	Names             []Name          `json:"names"`
	Popularity        Popularity      `json:"popularity"`
	LastUpdatedTime   JSONTime        `json:"lastUpdatedTime"`
	Locales           []Locale        `json:"locales,omitempty"`
	AlternateNames    []AlternateName `json:"alternateNames,omitempty"`
	LanguageOfContent []string        `json:"languageOfContent,omitempty"`
	Artists           []TrackArtist   `json:"artists"`
	Albums            []TrackAlbum    `json:"albums"`
	// Deleted must be a pointer to a true boolean or nil, never a pointer to a true boolean
	// This is a restriction of the api
	Deleted *bool `json:"deleted,omitempty"`
}

type TrackArtist struct {
	Id             string          `json:"id"`
	Names          []Name          `json:"names"`
	AlternateNames []AlternateName `json:"alternateNames,omitempty"`
}

type TrackAlbum struct {
	Id             string          `json:"id"`
	Names          []Name          `json:"names"`
	AlternateNames []AlternateName `json:"alternateNames,omitempty"`
	ReleaseType    string          `json:"releaseType,omitempty"`
}

var _ MetaDataReceiver = (*TrackCatalog)(nil)

func (tc *TrackCatalog) AddMetaData(md MetaData, lastUpdate time.Time) {
	if tc.Entities == nil {
		tc.Entities = []TrackEntity{}
	}

	lexTrackId := fmt.Sprintf("track.%s", md.TrackID())

	matched := false
	for _, entity := range tc.Entities {
		if entity.Id == lexTrackId {
			matched = true
			if lastUpdate.After(time.Time(entity.LastUpdatedTime)) {
				entity.LastUpdatedTime = JSONTime(lastUpdate)
			}
		}
	}
	if !matched {
		ne := TrackEntity{
			Id:              lexTrackId,
			Names:           []Name{{Language: "en", Value: md.Track()}},
			Popularity:      Popularity{Default: 100},
			LastUpdatedTime: JSONTime(lastUpdate),
			Locales:         DefaultLocales(),
		}
		ne.Artists = []TrackArtist{{
			Id:    fmt.Sprintf("artist.%s", md.ArtistID()),
			Names: []Name{{Language: "en", Value: md.Artist()}},
		}}
		ne.Albums = []TrackAlbum{{
			Id:    fmt.Sprintf("album.%s", md.AlbumID()),
			Names: []Name{{Language: "en", Value: md.Album()}},
		}}

		tc.Entities = append(tc.Entities, ne)
	}

}
