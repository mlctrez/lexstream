package catalog

import (
	"fmt"
	catalog_ "github.com/mlctrez/lexstream/smapiv0/catalog"
	"time"
)

// TrackCatalog
// https://developer.amazon.com/en-US/docs/alexa/music-skills/catalog-reference.html#track
type TrackCatalog struct {
	Header
	Entities []*TrackEntity `json:"entities"`
}

type TrackEntity struct {
	Id                string          `json:"id"`
	Names             []Name          `json:"names,omitempty"`
	Popularity        *Popularity     `json:"popularity,omitempty"`
	LastUpdatedTime   JSONTime        `json:"lastUpdatedTime"`
	Locales           []Locale        `json:"locales,omitempty"`
	AlternateNames    []AlternateName `json:"alternateNames,omitempty"`
	LanguageOfContent []string        `json:"languageOfContent,omitempty"`
	Artists           []TrackArtist   `json:"artists,omitempty"`
	Albums            []TrackAlbum    `json:"albums,omitempty"`
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

func CreateTrackCatalog() *TrackCatalog {
	return &TrackCatalog{Header: buildHeader(MusicRecording), Entities: []*TrackEntity{}}
}

var _ MetaDataReceiver = (*TrackCatalog)(nil)

func (tc *TrackCatalog) AddMetaData(md MetaData, lastUpdate time.Time) {

	if md.TrackID() == "" {
		return
	}

	lexTrackId := fmt.Sprintf("track.%s", md.TrackID())

	if md.Deleted() {
		deleteFlag := true
		ae := &TrackEntity{Id: lexTrackId, LastUpdatedTime: JSONTime(lastUpdate), Deleted: &deleteFlag}
		for i, entity := range tc.Entities {
			if entity.Id == lexTrackId {
				tc.Entities[i] = ae
				return
			}
		}
		tc.Entities = append(tc.Entities, ae)
		return
	}

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
		ne := &TrackEntity{
			Id:              lexTrackId,
			Names:           []Name{{Language: "en", Value: md.Track()}},
			Popularity:      &Popularity{Default: 100},
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

var _ StagingCatalog = (*TrackCatalog)(nil)

func (tc *TrackCatalog) CatalogType() catalog_.CatalogType {
	return catalog_.CatalogType_AMAZONMusicRecording()
}

func (tc *TrackCatalog) Validate() error {
	expectedCatalog := CreateTrackCatalog()
	if !expectedCatalog.TypeMatches(tc.Header) {
		return fmt.Errorf("staged catalog type does not match %s", expectedCatalog.Type)
	}
	if len(tc.Entities) == 0 {
		return fmt.Errorf("catalog has no entities")
	}
	// TODO: validate full catalog spec
	return nil
}
