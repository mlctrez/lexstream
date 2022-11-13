package skill

/*
Action Action of a resource.
*/
type Action string

func Action_CREATE() Action {
	return "CREATE"
}
func Action_UPDATE() Action {
	return "UPDATE"
}
func Action_ASSOCIATE() Action {
	return "ASSOCIATE"
}
func Action_DISASSOCIATE() Action {
	return "DISASSOCIATE"
}
