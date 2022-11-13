package manifest

/*
GadgetSupport Specifies if gadget support is required/optional for this skill to work.
*/
type GadgetSupport string

func GadgetSupport_REQUIRED() GadgetSupport {
	return "REQUIRED"
}
func GadgetSupport_OPTIONAL() GadgetSupport {
	return "OPTIONAL"
}
