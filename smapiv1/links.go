package smapiv1

/*
Links Links for the API navigation.
*/
type Links struct {
	Self *Link `json:"self,omitempty"`
	Next *Link `json:"next,omitempty"`
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
