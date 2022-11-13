package manifest

type AlexaForBusinessInterface struct {
	// Name of the interface.
	Namespace string   `json:"namespace,omitempty"`
	Version   *Version `json:"version,omitempty"`
	// Contains a list of requests/messages that skill can handle.
	Requests []*Request `json:"requests,omitempty"`
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
    "$ref": "#/definitions/v1.skill.Manifest.Request"
   },
   "type": "array"
  },
  "version": {
   "$ref": "#/definitions/v1.skill.Manifest.Version",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
