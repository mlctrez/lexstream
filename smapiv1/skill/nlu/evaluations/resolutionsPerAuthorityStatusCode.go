package evaluations

/*
ResolutionsPerAuthorityStatusCode A code indicating the results of attempting to resolve the user utterance against the defined slot types. This can be one of the following:
ER_SUCCESS_MATCH: The spoken value matched a value or synonym explicitly defined in your custom slot type. ER_SUCCESS_NO_MATCH: The spoken value did not match any values or synonyms explicitly defined in your custom slot type. ER_ERROR_TIMEOUT: An error occurred due to a timeout. ER_ERROR_EXCEPTION: An error occurred due to an exception during processing.
*/
type ResolutionsPerAuthorityStatusCode string

func ResolutionsPerAuthorityStatusCode_ER_SUCCESS_MATCH() ResolutionsPerAuthorityStatusCode {
	return "ER_SUCCESS_MATCH"
}
func ResolutionsPerAuthorityStatusCode_ER_SUCCESS_NO_MATCH() ResolutionsPerAuthorityStatusCode {
	return "ER_SUCCESS_NO_MATCH"
}
func ResolutionsPerAuthorityStatusCode_ER_ERROR_TIMEOUT() ResolutionsPerAuthorityStatusCode {
	return "ER_ERROR_TIMEOUT"
}
func ResolutionsPerAuthorityStatusCode_ER_ERROR_EXCEPTION() ResolutionsPerAuthorityStatusCode {
	return "ER_ERROR_EXCEPTION"
}
