package manifest

/*
SkillManifestCustomTask Defines the name and version of the task that the skill wants to handle.
*/
type SkillManifestCustomTask struct {
	// Name of the task.
	Name string `json:"name"`
	// Version of the task.
	Version string `json:"version"`
}

/*
{
 "description": "Defines the name and version of the task that the skill wants to handle.",
 "properties": {
  "name": {
   "description": "Name of the task.",
   "minLength": 1,
   "type": "string"
  },
  "version": {
   "description": "Version of the task.",
   "minLength": 1,
   "type": "string"
  }
 },
 "required": [
  "name",
  "version"
 ],
 "type": "object"
}
*/
