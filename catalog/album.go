package catalog

import (
	"fmt"
	catalog_ "github.com/mlctrez/lexstream/smapiv0/catalog"
	"time"
)

// AlbumCatalog
// https://developer.amazon.com/en-US/docs/alexa/music-skills/catalog-reference.html#album
type AlbumCatalog struct {
	Header
	Entities []*AlbumEntity `json:"entities"`
}

type AlbumEntity struct {
	Id                string          `json:"id"`
	Names             []Name          `json:"names,omitempty"`
	Popularity        *Popularity     `json:"popularity,omitempty"`
	LastUpdatedTime   JSONTime        `json:"lastUpdatedTime"`
	Locales           []Locale        `json:"locales,omitempty"`
	AlternateNames    []AlternateName `json:"alternateNames,omitempty"`
	LanguageOfContent []string        `json:"languageOfContent,omitempty"`
	ReleaseType       string          `json:"releaseType,omitempty"`
	Artists           []*AlbumArtist  `json:"artists,omitempty"`
	// Deleted must be nil ( not deleted ) or a pointer to true.
	//  deleted=false is not valid according to the catalog upload api
	Deleted *bool `json:"deleted,omitempty"`
}

type AlbumArtist struct {
	Id             string          `json:"id"`
	Names          []Name          `json:"names"`
	AlternateNames []AlternateName `json:"alternateNames,omitempty"`
}

func CreateAlbumCatalog() *AlbumCatalog {
	return &AlbumCatalog{Header: buildHeader(MusicAlbum), Entities: []*AlbumEntity{}}
}

var _ MetaDataReceiver = (*AlbumCatalog)(nil)

func (ac *AlbumCatalog) AddMetaData(md MetaData, lastUpdate time.Time) {

	if md.AlbumID() == "" {
		return
	}

	lexAlbumId := fmt.Sprintf("album.%s", md.AlbumID())

	var ae *AlbumEntity
	albumMatch := false

	if md.Deleted() {
		deleteFlag := true
		ae = &AlbumEntity{Id: lexAlbumId, LastUpdatedTime: JSONTime(lastUpdate), Deleted: &deleteFlag}
		for i, entity := range ac.Entities {
			if entity.Id == lexAlbumId {
				ac.Entities[i] = ae
				return
			}
		}
		ac.Entities = append(ac.Entities, ae)
		return
	}

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
			Popularity:      &Popularity{Default: 100},
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

var _ StagingCatalog = (*AlbumCatalog)(nil)

func (ac *AlbumCatalog) CatalogType() catalog_.CatalogType {
	return catalog_.CatalogType_AMAZONMusicAlbum()
}

func (ac *AlbumCatalog) Validate() error {
	expectedCatalog := CreateAlbumCatalog()
	if !expectedCatalog.TypeMatches(ac.Header) {
		return fmt.Errorf("staged catalog type does not match %s", expectedCatalog.Type)
	}
	if len(ac.Entities) == 0 {
		return fmt.Errorf("catalog has no entities")
	}
	// TODO: validate full catalog spec
	return nil
}
