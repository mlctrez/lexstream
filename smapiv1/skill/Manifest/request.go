package manifest

type Request struct {
	Name *RequestName `json:"name,omitempty"`
}

/*
{
 "properties": {
  "name": {
   "$ref": "#/definitions/v1.skill.Manifest.RequestName",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
