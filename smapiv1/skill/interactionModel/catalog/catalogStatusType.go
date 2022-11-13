package catalog

/*
CatalogStatusType Status of last modification request for a resource.
*/
type CatalogStatusType string

func CatalogStatusType_FAILED() CatalogStatusType {
	return "FAILED"
}
func CatalogStatusType_IN_PROGRESS() CatalogStatusType {
	return "IN_PROGRESS"
}
func CatalogStatusType_SUCCEEDED() CatalogStatusType {
	return "SUCCEEDED"
}
