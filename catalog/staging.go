package catalog

import (
	"fmt"
	"github.com/mlctrez/lexstream/internal/jutil"
	"github.com/mlctrez/lexstream/smapiv0/catalog"
	"path/filepath"
	"strings"
)

type StagingCatalog interface {
	CatalogType() catalog.CatalogType
	Validate() error
}

func StagingPathForType(catalogType catalog.CatalogType) string {
	prefixRemoved := strings.TrimPrefix(string(catalogType), "AMAZON.")
	name := fmt.Sprintf("%s.json", strings.ToLower(prefixRemoved))
	return filepath.Join("catalog_staging", name)
}

func WriteToStaging(catalog StagingCatalog) error {

	path := StagingPathForType(catalog.CatalogType())

	return jutil.WriteJson(path, catalog, true)
}

func ReadFromStaging(catalogType catalog.CatalogType) (any, error) {
	path := StagingPathForType(catalogType)
	target, err := NewCatalogForType(catalogType)
	if err != nil {
		return nil, err
	}

	err = jutil.DecodePath(path, target)
	return target, err
}

func NewCatalogForType(catalogType catalog.CatalogType) (any, error) {
	switch catalogType {
	case catalog.CatalogType_AMAZONMusicGroup():
		return CreateArtistCatalog(), nil
	case catalog.CatalogType_AMAZONMusicAlbum():
		return CreateAlbumCatalog(), nil
	case catalog.CatalogType_AMAZONMusicRecording():
		return CreateTrackCatalog(), nil
	case catalog.CatalogType_AMAZONMusicPlaylist():
		return CreatePlaylistCatalog(), nil
	default:
		return nil, fmt.Errorf("unsupported catalog type %s", catalogType)
	}
}
