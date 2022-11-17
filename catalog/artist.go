package catalog

import (
	"fmt"
	catalog_ "github.com/mlctrez/lexstream/smapiv0/catalog"
	"time"
)

// ArtistCatalog
// https://developer.amazon.com/en-US/docs/alexa/music-skills/catalog-reference.html#artist
type ArtistCatalog struct {
	Header
	Entities []*ArtistEntity `json:"entities"`
}

type ArtistEntity struct {
	Id                string          `json:"id"`
	Names             []Name          `json:"names,omitempty"`
	Popularity        *Popularity     `json:"popularity,omitempty"`
	LastUpdatedTime   JSONTime        `json:"lastUpdatedTime"`
	Locales           []Locale        `json:"locales,omitempty"`
	AlternateNames    []AlternateName `json:"alternateNames,omitempty"`
	LanguageOfContent []string        `json:"languageOfContent,omitempty"`
	// Deleted must be nil ( not deleted ) or a pointer to true.
	//  deleted=false is not valid according to the catalog upload api
	Deleted *bool `json:"deleted,omitempty"`
}

func CreateArtistCatalog() *ArtistCatalog {
	return &ArtistCatalog{Header: buildHeader(MusicGroup), Entities: []*ArtistEntity{}}
}

var _ MetaDataReceiver = (*ArtistCatalog)(nil)

func (ac *ArtistCatalog) AddMetaData(md MetaData, lastUpdate time.Time) {

	if md.ArtistID() == "" {
		return
	}

	lexArtistId := fmt.Sprintf("artist.%s", md.ArtistID())

	if md.Deleted() {
		deleteFlag := true
		ae := &ArtistEntity{Id: lexArtistId, LastUpdatedTime: JSONTime(lastUpdate), Deleted: &deleteFlag}
		for i, entity := range ac.Entities {
			if entity.Id == lexArtistId {
				ac.Entities[i] = ae
				return
			}
		}
		ac.Entities = append(ac.Entities, ae)
		return
	}

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
		ne := &ArtistEntity{
			Id:              lexArtistId,
			Names:           []Name{{Language: "en", Value: md.Artist()}},
			Popularity:      &Popularity{Default: 100},
			LastUpdatedTime: JSONTime(lastUpdate),
			Locales:         DefaultLocales(),
		}
		ac.Entities = append(ac.Entities, ne)
	}
}

var _ StagingCatalog = (*ArtistCatalog)(nil)

func (ac *ArtistCatalog) CatalogType() catalog_.CatalogType {
	return catalog_.CatalogType_AMAZONMusicGroup()
}

func (ac *ArtistCatalog) Validate() error {
	expectedCatalog := CreateArtistCatalog()
	if !expectedCatalog.TypeMatches(ac.Header) {
		return fmt.Errorf("staged catalog type does not match %s", expectedCatalog.Type)
	}
	if len(ac.Entities) == 0 {
		return fmt.Errorf("catalog has no entities")
	}
	// TODO: validate full catalog spec
	return nil

}
