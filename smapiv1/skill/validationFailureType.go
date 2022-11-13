package skill

/*
ValidationFailureType Enum for type of validation failure in the request.
*/
type ValidationFailureType string

func ValidationFailureType_RESOURCE_DOES_NOT_EXIST() ValidationFailureType {
	return "RESOURCE_DOES_NOT_EXIST"
}
func ValidationFailureType_RESOURCE_VERSION_DOES_NOT_MATCH() ValidationFailureType {
	return "RESOURCE_VERSION_DOES_NOT_MATCH"
}
func ValidationFailureType_MALFORMED_INPUT() ValidationFailureType {
	return "MALFORMED_INPUT"
}
func ValidationFailureType_EXPECTED_NOT_EMPTY_VALUE() ValidationFailureType {
	return "EXPECTED_NOT_EMPTY_VALUE"
}
func ValidationFailureType_INVALID_NUMBER_OF_OCCURENCES() ValidationFailureType {
	return "INVALID_NUMBER_OF_OCCURENCES"
}
func ValidationFailureType_INVALID_NUMBER_OF_PROPERTIES() ValidationFailureType {
	return "INVALID_NUMBER_OF_PROPERTIES"
}
func ValidationFailureType_EXPECTED_ATLEAST_ONE_RELATED_INSTANCE() ValidationFailureType {
	return "EXPECTED_ATLEAST_ONE_RELATED_INSTANCE"
}
func ValidationFailureType_EXPECTED_EXACTLY_ONE_RELATED_INSTANCE() ValidationFailureType {
	return "EXPECTED_EXACTLY_ONE_RELATED_INSTANCE"
}
func ValidationFailureType_RESOURCE_LOCKED() ValidationFailureType {
	return "RESOURCE_LOCKED"
}
func ValidationFailureType_UNEXPECTED_RESOURCE_STAGE() ValidationFailureType {
	return "UNEXPECTED_RESOURCE_STAGE"
}
func ValidationFailureType_UNEXPECTED_RESOURCE_PROPERTY() ValidationFailureType {
	return "UNEXPECTED_RESOURCE_PROPERTY"
}
func ValidationFailureType_MISSING_RESOURCE_PROPERTY() ValidationFailureType {
	return "MISSING_RESOURCE_PROPERTY"
}
