package evaluations

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

/*
Links Links for the API navigation.
*/
type Links struct {
	Self *smapiv1.Link `json:"self,omitempty"`
	Next *smapiv1.Link `json:"next,omitempty"`
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
