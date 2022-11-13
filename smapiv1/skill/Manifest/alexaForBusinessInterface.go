package manifest

type AlexaForBusinessInterface struct {
	// Name of the interface.
	Namespace string   `json,omitempty:"namespace"`
	Version   *Version `json,omitempty:"version"`
	// Contains a list of requests/messages that skill can handle.
	Requests []*Request `json,omitempty:"requests"`
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
