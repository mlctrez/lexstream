package skill

/*
ImageAttributes Set of properties of the image provided by the customer.
*/
type ImageAttributes struct {
	Dimension *ImageDimension `json:"dimension,omitempty"`
	Size      *ImageSize      `json:"size,omitempty"`
}

/*
{
 "description": "Set of properties of the image provided by the customer.",
 "properties": {
  "dimension": {
   "$ref": "#/definitions/v1.skill.ImageDimension"
  },
  "size": {
   "$ref": "#/definitions/v1.skill.ImageSize"
  }
 },
 "type": "object"
}
*/
