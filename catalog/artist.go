package catalog

import (
	"fmt"
	"time"
)

type ArtistCatalog struct {
	Header
	Entities []ArtistEntity `json:"entities"`
}

type ArtistEntity struct {
	Id                string          `json:"id"`
	Names             []Name          `json:"names"`
	Popularity        Popularity      `json:"popularity"`
	LastUpdatedTime   JSONTime        `json:"lastUpdatedTime"`
	Locales           []Locale        `json:"locales,omitempty"`
	AlternateNames    []AlternateName `json:"alternateNames,omitempty"`
	LanguageOfContent []string        `json:"languageOfContent,omitempty"`
	// Deleted must be nil ( not deleted ) or a pointer to true.
	//  deleted=false is not valid according to the catalog upload api
	Deleted *bool `json:"deleted,omitempty"`
}

func CreateArtistCatalog() *ArtistCatalog {
	return &ArtistCatalog{Header: buildHeader(MusicGroup)}
}

var _ MetaDataReceiver = (*ArtistCatalog)(nil)

func (ac *ArtistCatalog) AddMetaData(md MetaData, lastUpdate time.Time) {
	if ac.Entities == nil {
		ac.Entities = []ArtistEntity{}
	}

	lexArtistId := fmt.Sprintf("artist.%s", md.ArtistID())

	matched := false
	for _, entity := range ac.Entities {
		if entity.Id == lexArtistId {
			matched = true
			if lastUpdate.After(time.Time(entity.LastUpdatedTime)) {
				entity.LastUpdatedTime = JSONTime(lastUpdate)
			}
		}
	}
	if !matched {
		ne := ArtistEntity{
			Id:              lexArtistId,
			Names:           []Name{{Language: "en", Value: md.Artist()}},
			Popularity:      Popularity{Default: 100},
			LastUpdatedTime: JSONTime(lastUpdate),
			Locales:         DefaultLocales(),
		}
		ac.Entities = append(ac.Entities, ne)
	}
}
