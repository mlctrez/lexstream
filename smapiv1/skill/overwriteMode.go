package skill

/*
OverwriteMode Can locale of skill be overwritten. Default value is DO_NOT_OVERWRITE.
*/
type OverwriteMode string

func OverwriteMode_DO_NOT_OVERWRITE() OverwriteMode {
	return "DO_NOT_OVERWRITE"
}
func OverwriteMode_OVERWRITE() OverwriteMode {
	return "OVERWRITE"
}
