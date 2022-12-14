package manifest

type MusicInterfaces struct {
	// Name of the interface.
	Namespace string `json:"namespace,omitempty"`
	// Version of the interface.
	Version string `json:"version,omitempty"`
	// Contains a list of requests/messages that skill can handle.
	Requests []*MusicRequest `json:"requests,omitempty"`
}

/*
{
 "properties": {
  "namespace": {
   "description": "Name of the interface.",
   "type": "string"
  },
  "requests": {
   "description": "Contains a list of requests/messages that skill can handle.",
   "items": {
    "$ref": "#/definitions/v1.skill.Manifest.MusicRequest"
   },
   "type": "array"
  },
  "version": {
   "description": "Version of the interface.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
