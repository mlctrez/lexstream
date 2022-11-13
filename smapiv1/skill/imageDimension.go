package skill

/*
ImageDimension Dimensions of an image.
*/
type ImageDimension struct {
	// Width of the image in pixels.
	WidthInPixels int `json:"widthInPixels,omitempty"`
	// Height of the image in pixels.
	HeightInPixels int `json:"heightInPixels,omitempty"`
}

/*
{
 "description": "Dimensions of an image.",
 "properties": {
  "heightInPixels": {
   "description": "Height of the image in pixels.",
   "type": "integer"
  },
  "widthInPixels": {
   "description": "Width of the image in pixels.",
   "type": "integer"
  }
 },
 "type": "object"
}
*/
