package skill

type ValidationDataTypes string

func ValidationDataTypes_Object() ValidationDataTypes {
	return "object"
}
func ValidationDataTypes_Boolean() ValidationDataTypes {
	return "boolean"
}
func ValidationDataTypes_Integer() ValidationDataTypes {
	return "integer"
}
func ValidationDataTypes_Array() ValidationDataTypes {
	return "array"
}
func ValidationDataTypes_String() ValidationDataTypes {
	return "string"
}
func ValidationDataTypes_Null() ValidationDataTypes {
	return "null"
}
