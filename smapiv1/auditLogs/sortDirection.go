package auditlogs

/*
SortDirection Sets the sorting direction of the result items. When set to 'ASC' these items are returned in ascending order of sortField value and when set to 'DESC' these items are returned in descending order of sortField value.
*/
type SortDirection string

func SortDirection_ASC() SortDirection {
	return "ASC"
}
func SortDirection_DESC() SortDirection {
	return "DESC"
}
