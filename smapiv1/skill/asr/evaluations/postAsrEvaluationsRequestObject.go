package evaluations

type PostAsrEvaluationsRequestObject struct {
	Skill *Skill `json,omitempty:"skill"`
	// ID for annotation set
	AnnotationSetId string `json,omitempty:"annotationSetId"`
}

/*
{
 "properties": {
  "annotationSetId": {
   "description": "ID for annotation set",
   "type": "string"
  },
  "skill": {
   "$ref": "#/definitions/v1.skill.asr.evaluations.Skill"
  }
 },
 "required": [
  "annotationSetId",
  "skill"
 ],
 "type": "object"
}
*/
