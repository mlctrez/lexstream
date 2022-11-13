package skill

type ImportResponseSkill struct {
	SkillId   string                  `json,omitempty:"skillId"`
	ETag      string                  `json,omitempty:"eTag"`
	Resources []*ResourceImportStatus `json,omitempty:"resources"`
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
