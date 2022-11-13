package evaluations

type ResolutionsPerAuthorityStatus struct {
	/*
	   A code indicating the results of attempting to resolve the user utterance against the defined slot types. This can be one of the following:
	   ER_SUCCESS_MATCH: The spoken value matched a value or synonym explicitly defined in your custom slot type. ER_SUCCESS_NO_MATCH: The spoken value did not match any values or synonyms explicitly defined in your custom slot type. ER_ERROR_TIMEOUT: An error occurred due to a timeout. ER_ERROR_EXCEPTION: An error occurred due to an exception during processing.
	*/
	Code *ResolutionsPerAuthorityStatusCode `json:"code,omitempty"`
}

/*
{
 "properties": {
  "code": {
   "$ref": "#/definitions/v1.skill.nlu.evaluations.ResolutionsPerAuthorityStatusCode",
   "description": "A code indicating the results of attempting to resolve the user utterance against the defined slot types. This can be one of the following:\nER_SUCCESS_MATCH: The spoken value matched a value or synonym explicitly defined in your custom slot type. ER_SUCCESS_NO_MATCH: The spoken value did not match any values or synonyms explicitly defined in your custom slot type. ER_ERROR_TIMEOUT: An error occurred due to a timeout. ER_ERROR_EXCEPTION: An error occurred due to an exception during processing.\n",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
