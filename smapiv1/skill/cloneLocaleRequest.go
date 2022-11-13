package skill

/*
CloneLocaleRequest Defines the request body for the cloneLocale API.
*/
type CloneLocaleRequest struct {
	// Locale with the assets that will be cloned.
	SourceLocale string `json,omitempty:"sourceLocale"`
	// A list of locale(s) where the assets will be copied to.
	TargetLocales []string `json,omitempty:"targetLocales"`
	// Can locale of skill be overwritten. Default value is DO_NOT_OVERWRITE.
	OverwriteMode *OverwriteMode `json,omitempty:"overwriteMode"`
}

/*
{
 "description": "Defines the request body for the cloneLocale API.",
 "properties": {
  "overwriteMode": {
   "$ref": "#/definitions/v1.skill.OverwriteMode",
   "description": "Can locale of skill be overwritten. Default value is DO_NOT_OVERWRITE.",
   "x-isEnum": true
  },
  "sourceLocale": {
   "description": "Locale with the assets that will be cloned.",
   "type": "string"
  },
  "targetLocales": {
   "description": "A list of locale(s) where the assets will be copied to.",
   "items": {
    "type": "string"
   },
   "type": "array"
  }
 },
 "required": [
  "sourceLocale",
  "targetLocales"
 ],
 "type": "object"
}
*/
