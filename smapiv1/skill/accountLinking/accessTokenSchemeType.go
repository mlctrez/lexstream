package accountlinking

/*
AccessTokenSchemeType The type of client authentication scheme.
*/
type AccessTokenSchemeType string

func AccessTokenSchemeType_HTTP_BASIC() AccessTokenSchemeType {
	return "HTTP_BASIC"
}
func AccessTokenSchemeType_REQUEST_BODY_CREDENTIALS() AccessTokenSchemeType {
	return "REQUEST_BODY_CREDENTIALS"
}
