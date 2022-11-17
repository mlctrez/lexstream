package catalog

import "testing"

func TestHeader_TypeMatches(t *testing.T) {
	CreateAlbumCatalog().TypeMatches(Header{Type: string(MusicAlbum)})
}
