package skill

type UpdateSkillWithPackageRequest struct {
	// The URL for the skill package.
	Location string `json,omitempty:"location"`
}

/*
{
 "properties": {
  "location": {
   "description": "The URL for the skill package.",
   "format": "uri",
   "type": "string"
  }
 },
 "required": [
  "location"
 ],
 "type": "object"
}
*/
