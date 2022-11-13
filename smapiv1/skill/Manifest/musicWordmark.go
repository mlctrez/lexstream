package manifest

type MusicWordmark struct {
	// Wordmark logo to be used by devices with displays.
	Uri string `json,omitempty:"uri"`
}

/*
{
 "properties": {
  "uri": {
   "description": "Wordmark logo to be used by devices with displays.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
