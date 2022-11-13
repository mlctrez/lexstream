package history

/*
InteractionType Indicates if the utterance was executed as a "ONE_SHOT" interaction or "MODAL" interaction.
* `ONE_SHOT`: The user invokes the skill and states their intent in a single phrase.
* `MODAL`: The user first invokes the skill and then states their intent.
*/
type InteractionType string

func InteractionType_ONE_SHOT() InteractionType {
	return "ONE_SHOT"
}
func InteractionType_MODAL() InteractionType {
	return "MODAL"
}
