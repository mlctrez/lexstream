package manifest

/*
SmartHomeProtocol Version of the Smart Home API. Default and recommended value is '3'. You may create a skill with version '2' for testing migration to version '3', but a skill submission using version '2' will not be certified.
*/
type SmartHomeProtocol string

func SmartHomeProtocol_N1() SmartHomeProtocol {
	return "1"
}
func SmartHomeProtocol_N2() SmartHomeProtocol {
	return "2"
}
func SmartHomeProtocol_N25() SmartHomeProtocol {
	return "2.5"
}
func SmartHomeProtocol_N29() SmartHomeProtocol {
	return "2.9"
}
func SmartHomeProtocol_N3() SmartHomeProtocol {
	return "3"
}
