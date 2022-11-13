package subscription

/*
Event Represents an event that the subscriber is interested in. The event is of the format AlexaDevelopmentEvent.OPERATION. You can use wildcard event 'AlexaDevelopmentEvent.All' for recieving all developer notifications listed below.
  - 'AlexaDevelopmentEvent.ManifestUpdate' - The event representing the status of the update request on the Manifest.
  - 'AlexaDevelopmentEvent.SkillPublish' -   The event representing the status of the skill publish process.
  - 'AlexaDevelopmentEvent.SkillCertification' -   The event represents if a skill has been certified or not.
  - 'AlexaDevelopmentEvent.InteractionModelUpdate' -   The event represents the status of an Interaction Model build for a particular locale.
  - 'AlexaDevelopmentEvent.All' - A wildcard event name that allows subscription to all the existing events. While using this, you must not specify any other event name. AlexaDevelopmentEvent.All avoids the need of specifying every development event name in order to receive all events pertaining to a vendor account. Similarly, it avoids the need of updating an existing subscription to be able to receive new events, whenever supproted by notification service. Test Subscriber API cannot use this wildcard. Please make sure that your code can gracefully handle new/previously unknown events, if you are using this wildcard.
*/
type Event string

func Event_AlexaDevelopmentEventManifestUpdate() Event {
	return "AlexaDevelopmentEvent.ManifestUpdate"
}
func Event_AlexaDevelopmentEventSkillPublish() Event {
	return "AlexaDevelopmentEvent.SkillPublish"
}
func Event_AlexaDevelopmentEventSkillCertification() Event {
	return "AlexaDevelopmentEvent.SkillCertification"
}
func Event_AlexaDevelopmentEventInteractionModelUpdate() Event {
	return "AlexaDevelopmentEvent.InteractionModelUpdate"
}
func Event_AlexaDevelopmentEventAll() Event {
	return "AlexaDevelopmentEvent.All"
}
