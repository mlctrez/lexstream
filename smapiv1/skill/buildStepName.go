package skill

/*
BuildStepName Name of the build step. Possible values -
* `DIALOG_MODEL_BUILD` - Build status for dialog model.
* `LANGUAGE_MODEL_QUICK_BUILD` - Build status for FST model.
* `LANGUAGE_MODEL_FULL_BUILD` - Build status for statistical model.
*/
type BuildStepName string

func BuildStepName_DIALOG_MODEL_BUILD() BuildStepName {
	return "DIALOG_MODEL_BUILD"
}
func BuildStepName_LANGUAGE_MODEL_QUICK_BUILD() BuildStepName {
	return "LANGUAGE_MODEL_QUICK_BUILD"
}
func BuildStepName_LANGUAGE_MODEL_FULL_BUILD() BuildStepName {
	return "LANGUAGE_MODEL_FULL_BUILD"
}
