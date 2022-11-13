package evaluations

/*
EvaluationResult evaluation detailed result
*/
type EvaluationResult struct {
	Status     *EvaluationResultStatus   `json,omitempty:"status"`
	Annotation *AnnotationWithAudioAsset `json,omitempty:"annotation"`
	Output     *EvaluationResultOutput   `json,omitempty:"output"`
	Error      *ErrorObject              `json,omitempty:"error"`
}

/*
{
 "description": "evaluation detailed result",
 "properties": {
  "annotation": {
   "$ref": "#/definitions/v1.skill.asr.evaluations.AnnotationWithAudioAsset"
  },
  "error": {
   "$ref": "#/definitions/v1.skill.asr.evaluations.ErrorObject"
  },
  "output": {
   "$ref": "#/definitions/v1.skill.asr.evaluations.EvaluationResultOutput"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.asr.evaluations.EvaluationResultStatus",
   "x-isEnum": true
  }
 },
 "required": [
  "annotation",
  "status"
 ],
 "type": "object"
}
*/
