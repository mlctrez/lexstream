package smapiv1

type Error struct{}

/*
{
 "properties": {
  "code": {
   "description": "Error code that maps to an error message. Developers with different locales should be able to lookup the error description based on this code.\n",
   "type": "string"
  },
  "message": {
   "description": "Readable description of error. If standardized, this is generated from the error code and validation details.",
   "type": "string"
  }
 },
 "required": [
  "message"
 ],
 "type": "object",
 "x-baseType": true,
 "x-concreteType": true
}
*/
