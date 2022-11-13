package interactionmodel

/*
PromptItemsType Prompt can be specified in different formats e.g. text, ssml.
*/
type PromptItemsType string

func PromptItemsType_SSML() PromptItemsType {
	return "SSML"
}
func PromptItemsType_PlainText() PromptItemsType {
	return "PlainText"
}
