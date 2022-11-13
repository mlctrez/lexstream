package catalog

import (
	"fmt"
	"time"
)

type AlbumCatalog struct {
	Header
	Entities []*AlbumEntity `json:"entities"`
}

type AlbumEntity struct {
	Id                string          `json:"id"`
	Names             []Name          `json:"names"`
	Popularity        Popularity      `json:"popularity"`
	LastUpdatedTime   JSONTime        `json:"lastUpdatedTime"`
	Locales           []Locale        `json:"locales,omitempty"`
	AlternateNames    []AlternateName `json:"alternateNames,omitempty"`
	LanguageOfContent []string        `json:"languageOfContent,omitempty"`
	ReleaseType       string          `json:"releaseType,omitempty"`
	Artists           []*AlbumArtist  `json:"artists"`
}

type AlbumArtist struct {
	Id             string          `json:"id"`
	Names          []Name          `json:"names"`
	AlternateNames []AlternateName `json:"alternateNames,omitempty"`
}

func CreateAlbumCatalog() *AlbumCatalog {
	return &AlbumCatalog{Header: buildHeader(MusicAlbum)}
}

var _ MetaDataReceiver = (*AlbumCatalog)(nil)

func (ac *AlbumCatalog) AddMetaData(md MetaData, lastUpdate time.Time) {
	if ac.Entities == nil {
		ac.Entities = []*AlbumEntity{}
	}

	lexAlbumId := fmt.Sprintf("album.%s", md.AlbumID())

	var ae *AlbumEntity

	albumMatch := false
	for _, entity := range ac.Entities {
		if entity.Id == lexAlbumId {
			albumMatch = true
			ae = entity
			if lastUpdate.After(time.Time(entity.LastUpdatedTime)) {
				entity.LastUpdatedTime = JSONTime(lastUpdate)
			}
			// add in additional artists
			artistMatch := false
			for _, artist := range ae.Artists {
				if artist.Id == fmt.Sprintf("artist.%s", md.ArtistID()) {
					artistMatch = true
				}
			}
			if !artistMatch {
				na := &AlbumArtist{
					Id:    fmt.Sprintf("artist.%s", md.ArtistID()),
					Names: []Name{{Language: "en", Value: md.Artist()}},
				}
				ae.Artists = append(ae.Artists, na)
			}
		}
	}
	if !albumMatch {
		ae = &AlbumEntity{
			Id:              lexAlbumId,
			Names:           []Name{{Language: "en", Value: md.Album()}},
			Popularity:      Popularity{Default: 100},
			LastUpdatedTime: JSONTime(lastUpdate),
			Locales:         DefaultLocales(),
		}
		ae.Artists = []*AlbumArtist{{
			Id:    fmt.Sprintf("artist.%s", md.ArtistID()),
			Names: []Name{{Language: "en", Value: md.Artist()}},
		}}
		ac.Entities = append(ac.Entities, ae)
	}
}
