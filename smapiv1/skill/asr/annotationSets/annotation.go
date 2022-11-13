package annotationsets

/*
Annotation A single test case that describes the audio reference, expected transcriptions, test case weight etc. Each annotation object must have at least expectedTranscription or, uploadId and filePathInUpload in pair. In any case, filePathInUpload and uploadId must be present or missing in pair.
*/
type Annotation struct {
	// Upload id obtained when developer creates an upload using catalog API. Required to be present when expectedTranscription is missing. When uploadId is present, filePathInUpload must also be present.
	UploadId string `json:"uploadId,omitempty"`
	// File path in the uploaded zip file. For example, a zip containing a folder named 'a' and there is an audio b.mp3 in that folder. The file path is a/b.mp3. Notice that forward slash ('/') should be used to concatenate directories. Required to be present when expectedTranscription is missing. When filePathInUpload is present, uploadId must also be present.
	FilePathInUpload string `json:"filePathInUpload,omitempty"`
	/*
	   Weight of the test case in an evaluation, the value would be used for calculating metrics such as overall error rate.  The acceptable values are from 1 - 1000. 1 means least significant, 1000 means mot significant. Here is how weight  impact the `OVERALL_ERROR_RATE` calculation: For example, an annotation set consists of 3 annotations and they have weight of 8, 1, 1. The evaluation results  show that only the first annotation test case passed while the rest of the test cases failed. In this case,  the overall error rate is (8 / (8 + 1 + 1)) = 0.8
	*/
	EvaluationWeight int `json:"evaluationWeight,omitempty"`
	// Expected transcription text for the input audio. The acceptable length of the string is between 1 and 500 Unicode characters. Required to be present when uploadId and filePathInUpload are missing.
	ExpectedTranscription string `json:"expectedTranscription,omitempty"`
}

/*
{
 "description": "A single test case that describes the audio reference, expected transcriptions, test case weight etc. Each annotation object must have at least expectedTranscription or, uploadId and filePathInUpload in pair. In any case, filePathInUpload and uploadId must be present or missing in pair.",
 "properties": {
  "evaluationWeight": {
   "description": "Weight of the test case in an evaluation, the value would be used for calculating metrics such as overall error rate.  The acceptable values are from 1 - 1000. 1 means least significant, 1000 means mot significant. Here is how weight  impact the `OVERALL_ERROR_RATE` calculation: For example, an annotation set consists of 3 annotations and they have weight of 8, 1, 1. The evaluation results  show that only the first annotation test case passed while the rest of the test cases failed. In this case,  the overall error rate is (8 / (8 + 1 + 1)) = 0.8\n",
   "maximum": 1000,
   "minimum": 1,
   "type": "number"
  },
  "expectedTranscription": {
   "description": "Expected transcription text for the input audio. The acceptable length of the string is between 1 and 500 Unicode characters. Required to be present when uploadId and filePathInUpload are missing.",
   "type": "string"
  },
  "filePathInUpload": {
   "description": "File path in the uploaded zip file. For example, a zip containing a folder named 'a' and there is an audio b.mp3 in that folder. The file path is a/b.mp3. Notice that forward slash ('/') should be used to concatenate directories. Required to be present when expectedTranscription is missing. When filePathInUpload is present, uploadId must also be present.",
   "type": "string"
  },
  "uploadId": {
   "description": "Upload id obtained when developer creates an upload using catalog API. Required to be present when expectedTranscription is missing. When uploadId is present, filePathInUpload must also be present.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
