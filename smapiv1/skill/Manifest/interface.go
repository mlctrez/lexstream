package manifest

type Interface struct {
	Type_ string `json:"type"`
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
