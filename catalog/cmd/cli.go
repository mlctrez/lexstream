package main

import (
	"github.com/mlctrez/lexstream/catalog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	builder := catalog.New()

	//builder.AddMetaData(catalog.DeletedMetaData("001", "001", "001", ""), time.Now())
	//builder.ScanUserMusicFolder()

	builder.AddMetaData(&catalog.StaticMetaData{
		ArtistValue:     "Toby Fox",
		ArtistIDValue:   "001",
		AlbumValue:      "Undertale Soundtrack",
		AlbumIDValue:    "001",
		TrackValue:      "Home",
		TrackIDValue:    "001",
		PlaylistValue:   "Current",
		PlaylistIDValue: "001",
		DeletedValue:    false,
	}, time.Now())

	dir := "catalog_staging"
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		panic(err)
	}

	if err = catalog.WriteToStaging(builder.ArtistCatalog); err != nil {
		panic(err)
	}

	if err = catalog.WriteToStaging(builder.AlbumCatalog); err != nil {
		panic(err)
	}
	if err = catalog.WriteToStaging(builder.TrackCatalog); err != nil {
		panic(err)
	}
	if err = catalog.WriteToStaging(builder.PlaylistCatalog); err != nil {
		panic(err)
	}

}

func jsFile(dir string, in catalog.HeaderType) string {
	noPrefix := strings.TrimPrefix(string(in), "AMAZON.")
	name := strings.ToLower(noPrefix + ".json")
	return filepath.Join(dir, name)
}
