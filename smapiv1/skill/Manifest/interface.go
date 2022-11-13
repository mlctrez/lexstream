package manifest

type Interface struct {
	Type_ string `json:"type,omitempty"`
}

/*
{
 "discriminator": "type",
 "properties": {
  "type": {
   "type": "string",
   "x-isDiscriminator": true
  }
 },
 "required": [
  "type"
 ],
 "type": "object"
}
*/
