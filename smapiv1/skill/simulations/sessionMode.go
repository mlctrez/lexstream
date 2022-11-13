package simulations

/*
SessionMode Indicate the session mode of the current simulation is using.
*/
type SessionMode string

func SessionMode_DEFAULT() SessionMode {
	return "DEFAULT"
}
func SessionMode_FORCE_NEW_SESSION() SessionMode {
	return "FORCE_NEW_SESSION"
}
