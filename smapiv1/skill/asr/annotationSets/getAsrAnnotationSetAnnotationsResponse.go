package annotationsets

/*
GetAsrAnnotationSetAnnotationsResponse This is the payload schema for annotation set contents. Note that when uploadId and filePathInUpload is present, and the payload content type is 'application/json', audioAsset is included in the returned annotation set content payload. For 'text/csv' annotation set content type, audioAssetDownloadUrl and audioAssetDownloadUrlExpiryTime are included in the csv headers for representing the audio download url and the expiry time of the presigned audio download.
*/
type GetAsrAnnotationSetAnnotationsResponse struct {
	Annotations       []*AnnotationWithAudioAsset `json:"annotations,omitempty"`
	PaginationContext *PaginationContext          `json:"paginationContext,omitempty"`
}

/*
{
 "description": "This is the payload schema for annotation set contents. Note that when uploadId and filePathInUpload is present, and the payload content type is 'application/json', audioAsset is included in the returned annotation set content payload. For 'text/csv' annotation set content type, audioAssetDownloadUrl and audioAssetDownloadUrlExpiryTime are included in the csv headers for representing the audio download url and the expiry time of the presigned audio download.\n",
 "properties": {
  "annotations": {
   "items": {
    "$ref": "#/definitions/v1.skill.asr.annotationSets.AnnotationWithAudioAsset"
   },
   "type": "array"
  },
  "paginationContext": {
   "$ref": "#/definitions/v1.skill.asr.annotationSets.PaginationContext"
  }
 },
 "required": [
  "annotations"
 ],
 "type": "object"
}
*/
