package smapiv1

/*
Links Links for the API navigation.
*/
type Links struct {
	Self *Link `json,omitempty:"self"`
	Next *Link `json,omitempty:"next"`
}

/*
{
 "description": "Links for the API navigation.",
 "properties": {
  "next": {
   "$ref": "#/definitions/v1.Link"
  },
  "self": {
   "$ref": "#/definitions/v1.Link"
  }
 },
 "type": "object"
}
*/
