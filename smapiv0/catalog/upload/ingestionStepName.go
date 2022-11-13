package upload

type IngestionStepName string

func IngestionStepName_UPLOAD() IngestionStepName {
	return "UPLOAD"
}
func IngestionStepName_SCHEMA_VALIDATION() IngestionStepName {
	return "SCHEMA_VALIDATION"
}
