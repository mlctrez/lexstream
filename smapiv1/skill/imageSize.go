package skill

/*
ImageSize On disk storage size of image.
*/
type ImageSize struct {
	// Value measuring the size of the image.
	Value int            `json:"value"`
	Unit  *ImageSizeUnit `json:"unit"`
}

/*
{
 "description": "On disk storage size of image.",
 "properties": {
  "unit": {
   "$ref": "#/definitions/v1.skill.ImageSizeUnit",
   "x-isEnum": true
  },
  "value": {
   "description": "Value measuring the size of the image.",
   "type": "number"
  }
 },
 "type": "object"
}
*/
