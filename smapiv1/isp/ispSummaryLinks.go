package isp

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

type IspSummaryLinks struct {
	Self *smapiv1.Link `json:"self,omitempty"`
}

/*
{
 "properties": {
  "self": {
   "$ref": "#/definitions/v1.Link"
  }
 },
 "type": "object"
}
*/
