package smapiv0

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
   "$ref": "#/definitions/v0.Link"
  },
  "self": {
   "$ref": "#/definitions/v0.Link"
  }
 },
 "type": "object"
}
*/
