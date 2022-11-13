package auditlogs

/*
SortField Sets the field on which the sorting would be applied.
*/
type SortField string

func SortField_Timestamp() SortField {
	return "timestamp"
}
func SortField_Operation() SortField {
	return "operation"
}
func SortField_Resourceid() SortField {
	return "resource.id"
}
func SortField_Resourcetype() SortField {
	return "resource.type"
}
func SortField_RequesteruserId() SortField {
	return "requester.userId"
}
func SortField_Clientid() SortField {
	return "client.id"
}
func SortField_HttpResponseCode() SortField {
	return "httpResponseCode"
}
