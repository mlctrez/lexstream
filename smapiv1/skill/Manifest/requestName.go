package manifest

/*
RequestName Name of the request.
*/
type RequestName string

func RequestName_Search() RequestName {
	return "Search"
}
func RequestName_Create() RequestName {
	return "Create"
}
func RequestName_Update() RequestName {
	return "Update"
}
