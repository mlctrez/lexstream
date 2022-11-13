package accountlinking

/*
AccountLinkingType The type of account linking.
*/
type AccountLinkingType string

func AccountLinkingType_AUTH_CODE() AccountLinkingType {
	return "AUTH_CODE"
}
func AccountLinkingType_IMPLICIT() AccountLinkingType {
	return "IMPLICIT"
}
