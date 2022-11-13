package validations

/*
ResponseValidationImportance String that specifies importance of the validation. Possible values are "REQUIRED" and "RECOMMENDED"
*/
type ResponseValidationImportance string

func ResponseValidationImportance_REQUIRED() ResponseValidationImportance {
	return "REQUIRED"
}
func ResponseValidationImportance_RECOMMENDED() ResponseValidationImportance {
	return "RECOMMENDED"
}
