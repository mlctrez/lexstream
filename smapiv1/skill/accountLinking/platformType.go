package accountlinking

/*
PlatformType Defines the type of platform that will be used by the customer to perform account linking.
*/
type PlatformType string

func PlatformType_IOS() PlatformType {
	return "iOS"
}
func PlatformType_Android() PlatformType {
	return "Android"
}
