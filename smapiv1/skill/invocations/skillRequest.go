package invocations

type SkillRequest struct {
	/*
	   ASK request body schema as defined in the public facing documentation (https://tiny.amazon.com/1h8keglep/deveamazpublsolualexalexdocs)
	*/
	Body map[string]any `json,omitempty:"body"`
}

/*
{
 "properties": {
  "body": {
   "description": "ASK request body schema as defined in the public facing documentation (https://tiny.amazon.com/1h8keglep/deveamazpublsolualexalexdocs)\n",
   "properties": {},
   "type": "object"
  }
 },
 "required": [
  "body"
 ],
 "type": "object"
}
*/
