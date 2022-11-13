package amsapi

// Art contains artwork for a media content. Images must be in PNG or JPG file format.
// For other requirements for Alexa skill content, including background images, see Policy Testing.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#art
type Art struct {
	Sources []ArtSource `json:"sources"`
}

// ArtSource contains information about a single size of a media content's art (for example, album cover art).
// Images must be in PNG or JPG file format.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#artsource
type ArtSource struct {
	// The URL for the image of this size. URL must be HTTPS.
	Url string `json:"url"`
	// The size for the image. Alexa can use this optional field to help determine which size of ArtSource to use.
	//	"X_SMALL" (recommended for images approximately 48 x 48 pixels)
	//	"SMALL"   (recommended for 60 x 60 pixels)
	//	"MEDIUM"  (recommended for 110 x 110 pixels)
	//	"LARGE"   (recommended for 256 x 256 pixels)
	//	"X_LARGE" (recommended for 600 x 600 pixels)
	Size         string `json:"size,omitempty"`
	WidthPixels  int    `json:"widthPixels,omitempty"`
	HeightPixels int    `json:"heightPixels,omitempty"`
}
