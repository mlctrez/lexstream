package evaluations

/*
AnnotationWithAudioAsset object containing annotation content and audio file download information.
*/
type AnnotationWithAudioAsset struct {
	Annotation
	AudioAsset *AudioAsset `json:"audioAsset,omitempty"`
}

/*
{
 "properties": {
  "audioAsset": {
   "$ref": "#/definitions/v1.skill.asr.evaluations.AudioAsset"
  }
 },
 "required": [
  "audioAsset"
 ],
 "type": "object"
}
*/
