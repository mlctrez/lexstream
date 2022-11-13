package main

import (
	"encoding/json"
	"github.com/mlctrez/lexstream/catalog"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	builder := catalog.New()

	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	err = builder.Walk(filepath.Join(dir, "Music"))
	if err != nil {
		panic(err)
	}

	dir = "catalog_staging"
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		panic(err)
	}

	jsFile := func(in catalog.HeaderType) string {
		noPrefix := strings.TrimPrefix(string(in), "AMAZON.")
		return strings.ToLower(noPrefix + ".json")
	}

	writeCatalogFile(filepath.Join(dir, jsFile(catalog.MusicGroup)), builder.ArtistCatalog)
	writeCatalogFile(filepath.Join(dir, jsFile(catalog.MusicAlbum)), builder.AlbumCatalog)
	writeCatalogFile(filepath.Join(dir, jsFile(catalog.MusicRecording)), builder.TrackCatalog)

}

func writeCatalogFile(path string, i interface{}) {
	out, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	writeCatalog(out, i)
}

func writeCatalog(out io.WriteCloser, i interface{}) {
	defer func() { _ = out.Close() }()
	err := json.NewEncoder(out).Encode(i)
	if err != nil {
		panic(err)
	}
}
