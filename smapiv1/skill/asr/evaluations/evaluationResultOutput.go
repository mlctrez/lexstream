package evaluations

type EvaluationResultOutput struct {
	// actual transcription returned from ASR for the audio
	Transcription string `json:"transcription"`
}

/*
{
 "properties": {
  "transcription": {
   "description": "actual transcription returned from ASR for the audio",
   "type": "string"
  }
 },
 "required": [
  "transcription"
 ],
 "type": "object"
}
*/
