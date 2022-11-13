package manifest

type MusicCapability struct {
	// Namespace of music skill api.
	Namespace string `json:"namespace"`
	// Name of music skill api.
	Name string `json:"name"`
	// Version of music skill api.
	Version string `json:"version"`
}

/*
{
 "properties": {
  "name": {
   "description": "Name of music skill api.",
   "type": "string"
  },
  "namespace": {
   "description": "Namespace of music skill api.",
   "type": "string"
  },
  "version": {
   "description": "Version of music skill api.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
