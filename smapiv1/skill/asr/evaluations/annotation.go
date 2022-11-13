package evaluations

type Annotation struct {
	// upload id obtained when developer creates an upload using catalog API
	UploadId string `json:"uploadId,omitempty"`
	// file path in the uploaded zip file. For example, a zip containing a folder named 'a' and there is an audio b.mp3 in that folder. The file path is a/b.mp3. Notice that forward slash ('/') should be used to concatenate directories.
	FilePathInUpload string `json:"filePathInUpload,omitempty"`
	/*
	   weight of the test case in an evaluation, the value would be used for calculating metrics such as overall error rate.  The acceptable values are from 1 - 1000. 1 means least significant, 1000 means mot significant. Here is how weight  impact the `OVERALL_ERROR_RATE` calculation: For example, an annotation set consists of 3 annotations and they have weight of 8, 1, 1. The evaluation results  show that only the first annotation test case passed while the rest of the test cases failed. In this case,  the overall error rate is (8 / (8 + 1 + 1)) = 0.8
	*/
	EvaluationWeight int `json:"evaluationWeight,omitempty"`
	// expected transcription text for the input audio. The acceptable length of the string is between 1 and 500 Unicode characters
	ExpectedTranscription string `json:"expectedTranscription,omitempty"`
}

/*
{
 "properties": {
  "evaluationWeight": {
   "description": "weight of the test case in an evaluation, the value would be used for calculating metrics such as overall error rate.  The acceptable values are from 1 - 1000. 1 means least significant, 1000 means mot significant. Here is how weight  impact the `OVERALL_ERROR_RATE` calculation: For example, an annotation set consists of 3 annotations and they have weight of 8, 1, 1. The evaluation results  show that only the first annotation test case passed while the rest of the test cases failed. In this case,  the overall error rate is (8 / (8 + 1 + 1)) = 0.8\n",
   "maximum": 1000,
   "minimum": 1,
   "type": "number"
  },
  "expectedTranscription": {
   "description": "expected transcription text for the input audio. The acceptable length of the string is between 1 and 500 Unicode characters",
   "type": "string"
  },
  "filePathInUpload": {
   "description": "file path in the uploaded zip file. For example, a zip containing a folder named 'a' and there is an audio b.mp3 in that folder. The file path is a/b.mp3. Notice that forward slash ('/') should be used to concatenate directories.",
   "type": "string"
  },
  "uploadId": {
   "description": "upload id obtained when developer creates an upload using catalog API",
   "type": "string"
  }
 },
 "required": [
  "evaluationWeight",
  "expectedTranscription",
  "filePathInUpload",
  "uploadId"
 ],
 "type": "object"
}
*/
