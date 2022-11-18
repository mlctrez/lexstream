package coverart

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const ArchiveUrl = "https://coverartarchive.org/release/%s"

// FromAlbumId retrieves cover art for the provided musicbrainz release id.
func FromAlbumId(id string) (coverArt *CoverArt, err error) {

	request, err := http.NewRequest("GET", fmt.Sprintf(ArchiveUrl, id), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	coverArt = &CoverArt{}
	err = json.NewDecoder(response.Body).Decode(coverArt)
	return
}

// FrontLarge returns the large thumbnail url for the front cover.
func (ca *CoverArt) FrontLarge() (thumbUrl string, err error) {
	for _, image := range ca.Images {
		if image.Front {
			// "500" is the new format, "large" is legacy. look for both
			for _, size := range []string{"500", "large"} {
				if tn, ok := image.Thumbnails[size]; ok {
					var tnUrl *url.URL
					tnUrl, err = url.Parse(tn)
					if err != nil {
						return
					}
					// correct url scheme to make alexa happy
					tnUrl.Scheme = "https"
					thumbUrl = tnUrl.String()
					return
				}
			}
		}
	}
	err = fmt.Errorf("thumbnail not found")
	return
}

type CoverArt struct {
	Images  []Image `json:"images"`
	Release string  `json:"release"`
}

type Image struct {
	Image      string            `json:"image"`
	Edit       int               `json:"edit"`
	Comment    string            `json:"comment"`
	Thumbnails map[string]string `json:"thumbnails"`
	Id         int64             `json:"id"`
	Types      []string          `json:"types"`
	Approved   bool              `json:"approved"`
	Back       bool              `json:"back"`
	Front      bool              `json:"front"`
}
