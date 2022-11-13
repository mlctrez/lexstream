package annotationsets

/*
AnnotationWithAudioAsset Object containing annotation content and audio file download information.
*/
type AnnotationWithAudioAsset struct {
	Annotation
	AudioAsset *AudioAsset `json:"audioAsset"`
}

/*
{
 "properties": {
  "audioAsset": {
   "$ref": "#/definitions/v1.skill.asr.annotationSets.AudioAsset"
  }
 },
 "type": "object"
}
*/
