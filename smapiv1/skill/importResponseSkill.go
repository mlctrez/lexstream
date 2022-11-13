package skill

type ImportResponseSkill struct {
	SkillId   string                  `json:"skillId"`
	ETag      string                  `json:"eTag"`
	Resources []*ResourceImportStatus `json:"resources"`
}

/*
{
 "properties": {
  "eTag": {
   "type": "string"
  },
  "resources": {
   "items": {
    "$ref": "#/definitions/v1.skill.ResourceImportStatus"
   },
   "type": "array"
  },
  "skillId": {
   "type": "string"
  }
 },
 "required": [
  "resources"
 ],
 "type": "object"
}
*/
