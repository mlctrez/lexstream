package catalog

import (
	"fmt"
	catalog_ "github.com/mlctrez/lexstream/smapiv0/catalog"
	"time"
)

// PlaylistCatalog
// https://developer.amazon.com/en-US/docs/alexa/music-skills/catalog-reference.html#playlist
type PlaylistCatalog struct {
	Header
	Entities []*PlaylistEntity `json:"entities"`
}

type PlaylistEntity struct {
	Id              string          `json:"id"`
	Names           []Name          `json:"names"`
	Popularity      Popularity      `json:"popularity"`
	LastUpdatedTime JSONTime        `json:"lastUpdatedTime"`
	Locales         []Locale        `json:"locales,omitempty"`
	AlternateNames  []AlternateName `json:"alternateNames,omitempty"`
	// Deleted must be nil ( not deleted ) or a pointer to true.
	//  deleted=false is not valid according to the catalog upload api
	Deleted *bool `json:"deleted,omitempty"`
}

func CreatePlaylistCatalog() *PlaylistCatalog {
	return &PlaylistCatalog{Header: buildHeader(MusicPlaylist), Entities: []*PlaylistEntity{}}
}

var _ MetaDataReceiver = (*PlaylistCatalog)(nil)

func (pc *PlaylistCatalog) AddMetaData(md MetaData, lastUpdate time.Time) {
	if md.PlaylistID() == "" {
		return
	}
	lexPlaylistID := fmt.Sprintf("playlist.%s", md.PlaylistID())

	if md.Deleted() {
		deleteFlag := true
		ae := &PlaylistEntity{Id: lexPlaylistID, LastUpdatedTime: JSONTime(lastUpdate), Deleted: &deleteFlag}
		for i, entity := range pc.Entities {
			if entity.Id == lexPlaylistID {
				pc.Entities[i] = ae
				return
			}
		}
		pc.Entities = append(pc.Entities, ae)
		return
	}

	var pe *PlaylistEntity
	for _, entity := range pc.Entities {
		if entity.Id == lexPlaylistID {
			pe = entity
		}
	}
	if pe == nil {
		pe = &PlaylistEntity{
			Id:              fmt.Sprintf("playlist.%s", md.PlaylistID()),
			Names:           []Name{{Language: "en", Value: md.Playlist()}},
			Popularity:      Popularity{Default: 100},
			LastUpdatedTime: JSONTime(lastUpdate),
			Locales:         DefaultLocales(),
		}
		pc.Entities = append(pc.Entities, pe)
	} else {
		if lastUpdate.After(time.Time(pe.LastUpdatedTime)) {
			pe.LastUpdatedTime = JSONTime(lastUpdate)
		}
		pe.Names = []Name{{Language: "en", Value: md.Playlist()}}
	}

}

var _ StagingCatalog = (*PlaylistCatalog)(nil)

func (pc *PlaylistCatalog) CatalogType() catalog_.CatalogType {
	return catalog_.CatalogType_AMAZONMusicPlaylist()
}

func (pc *PlaylistCatalog) Validate() error {
	expectedCatalog := CreatePlaylistCatalog()
	if !expectedCatalog.TypeMatches(pc.Header) {
		return fmt.Errorf("staged catalog type does not match %s", expectedCatalog.Type)
	}
	if len(pc.Entities) == 0 {
		return fmt.Errorf("catalog has no entities")
	}
	// TODO: validate full catalog spec
	return nil
}
