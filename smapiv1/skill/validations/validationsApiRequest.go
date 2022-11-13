package validations

type ValidationsApiRequest struct {
	Locales []string `json:"locales,omitempty"`
}

/*
{
 "properties": {
  "locales": {
   "items": {
    "type": "string"
   },
   "type": "array"
  }
 },
 "required": [
  "locales"
 ],
 "type": "object"
}
*/
