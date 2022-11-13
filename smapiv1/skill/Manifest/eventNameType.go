package manifest

/*
EventNameType Name of the event to be subscribed to.
*/
type EventNameType string

func EventNameType_SKILL_ENABLED() EventNameType {
	return "SKILL_ENABLED"
}
func EventNameType_SKILL_DISABLED() EventNameType {
	return "SKILL_DISABLED"
}
func EventNameType_SKILL_PERMISSION_ACCEPTED() EventNameType {
	return "SKILL_PERMISSION_ACCEPTED"
}
func EventNameType_SKILL_PERMISSION_CHANGED() EventNameType {
	return "SKILL_PERMISSION_CHANGED"
}
func EventNameType_SKILL_ACCOUNT_LINKED() EventNameType {
	return "SKILL_ACCOUNT_LINKED"
}
func EventNameType_ITEMS_CREATED() EventNameType {
	return "ITEMS_CREATED"
}
func EventNameType_ITEMS_UPDATED() EventNameType {
	return "ITEMS_UPDATED"
}
func EventNameType_ITEMS_DELETED() EventNameType {
	return "ITEMS_DELETED"
}
func EventNameType_LIST_CREATED() EventNameType {
	return "LIST_CREATED"
}
func EventNameType_LIST_UPDATED() EventNameType {
	return "LIST_UPDATED"
}
func EventNameType_LIST_DELETED() EventNameType {
	return "LIST_DELETED"
}
func EventNameType_ALL_LISTS_CHANGED() EventNameType {
	return "ALL_LISTS_CHANGED"
}
func EventNameType_REMINDER_STARTED() EventNameType {
	return "REMINDER_STARTED"
}
func EventNameType_REMINDER_CREATED() EventNameType {
	return "REMINDER_CREATED"
}
func EventNameType_REMINDER_UPDATED() EventNameType {
	return "REMINDER_UPDATED"
}
func EventNameType_REMINDER_DELETED() EventNameType {
	return "REMINDER_DELETED"
}
func EventNameType_REMINDER_STATUS_CHANGED() EventNameType {
	return "REMINDER_STATUS_CHANGED"
}
func EventNameType_AUDIO_ITEM_PLAYBACK_STARTED() EventNameType {
	return "AUDIO_ITEM_PLAYBACK_STARTED"
}
func EventNameType_AUDIO_ITEM_PLAYBACK_FINISHED() EventNameType {
	return "AUDIO_ITEM_PLAYBACK_FINISHED"
}
func EventNameType_AUDIO_ITEM_PLAYBACK_STOPPED() EventNameType {
	return "AUDIO_ITEM_PLAYBACK_STOPPED"
}
func EventNameType_AUDIO_ITEM_PLAYBACK_FAILED() EventNameType {
	return "AUDIO_ITEM_PLAYBACK_FAILED"
}
func EventNameType_SKILL_PROACTIVE_SUBSCRIPTION_CHANGED() EventNameType {
	return "SKILL_PROACTIVE_SUBSCRIPTION_CHANGED"
}
func EventNameType_IN_SKILL_PRODUCT_SUBSCRIPTION_STARTED() EventNameType {
	return "IN_SKILL_PRODUCT_SUBSCRIPTION_STARTED"
}
func EventNameType_IN_SKILL_PRODUCT_SUBSCRIPTION_RENEWED() EventNameType {
	return "IN_SKILL_PRODUCT_SUBSCRIPTION_RENEWED"
}
func EventNameType_IN_SKILL_PRODUCT_SUBSCRIPTION_ENDED() EventNameType {
	return "IN_SKILL_PRODUCT_SUBSCRIPTION_ENDED"
}
func EventNameType_LegacyActivityManagerActivityContextRemovedEvent() EventNameType {
	return "Legacy.ActivityManager.ActivityContextRemovedEvent"
}
func EventNameType_LegacyActivityManagerActivityInterrupted() EventNameType {
	return "Legacy.ActivityManager.ActivityInterrupted"
}
func EventNameType_LegacyActivityManagerFocusChanged() EventNameType {
	return "Legacy.ActivityManager.FocusChanged"
}
func EventNameType_LegacyAlertsControllerDismissCommand() EventNameType {
	return "Legacy.AlertsController.DismissCommand"
}
func EventNameType_LegacyAlertsControllerSnoozeCommand() EventNameType {
	return "Legacy.AlertsController.SnoozeCommand"
}
func EventNameType_LegacyAudioPlayerAudioStutter() EventNameType {
	return "Legacy.AudioPlayer.AudioStutter"
}
func EventNameType_LegacyAudioPlayerInitialPlaybackProgressReport() EventNameType {
	return "Legacy.AudioPlayer.InitialPlaybackProgressReport"
}
func EventNameType_LegacyAudioPlayerMetadata() EventNameType {
	return "Legacy.AudioPlayer.Metadata"
}
func EventNameType_LegacyAudioPlayerPeriodicPlaybackProgressReport() EventNameType {
	return "Legacy.AudioPlayer.PeriodicPlaybackProgressReport"
}
func EventNameType_LegacyAudioPlayerPlaybackError() EventNameType {
	return "Legacy.AudioPlayer.PlaybackError"
}
func EventNameType_LegacyAudioPlayerPlaybackFinished() EventNameType {
	return "Legacy.AudioPlayer.PlaybackFinished"
}
func EventNameType_LegacyAudioPlayerPlaybackIdle() EventNameType {
	return "Legacy.AudioPlayer.PlaybackIdle"
}
func EventNameType_LegacyAudioPlayerPlaybackInterrupted() EventNameType {
	return "Legacy.AudioPlayer.PlaybackInterrupted"
}
func EventNameType_LegacyAudioPlayerPlaybackNearlyFinished() EventNameType {
	return "Legacy.AudioPlayer.PlaybackNearlyFinished"
}
func EventNameType_LegacyAudioPlayerPlaybackPaused() EventNameType {
	return "Legacy.AudioPlayer.PlaybackPaused"
}
func EventNameType_LegacyAudioPlayerPlaybackResumed() EventNameType {
	return "Legacy.AudioPlayer.PlaybackResumed"
}
func EventNameType_LegacyAudioPlayerPlaybackStarted() EventNameType {
	return "Legacy.AudioPlayer.PlaybackStarted"
}
func EventNameType_LegacyAudioPlayerPlaybackStutterFinished() EventNameType {
	return "Legacy.AudioPlayer.PlaybackStutterFinished"
}
func EventNameType_LegacyAudioPlayerPlaybackStutterStarted() EventNameType {
	return "Legacy.AudioPlayer.PlaybackStutterStarted"
}
func EventNameType_LegacyAudioPlayerGuiButtonClickedEvent() EventNameType {
	return "Legacy.AudioPlayerGui.ButtonClickedEvent"
}
func EventNameType_LegacyAudioPlayerGuiLyricsViewedEvent() EventNameType {
	return "Legacy.AudioPlayerGui.LyricsViewedEvent"
}
func EventNameType_LegacyAuxControllerDirectionChanged() EventNameType {
	return "Legacy.AuxController.DirectionChanged"
}
func EventNameType_LegacyAuxControllerEnabledStateChanged() EventNameType {
	return "Legacy.AuxController.EnabledStateChanged"
}
func EventNameType_LegacyAuxControllerInputActivityStateChanged() EventNameType {
	return "Legacy.AuxController.InputActivityStateChanged"
}
func EventNameType_LegacyAuxControllerPluggedStateChanged() EventNameType {
	return "Legacy.AuxController.PluggedStateChanged"
}
func EventNameType_LegacyBluetoothNetworkCancelPairingMode() EventNameType {
	return "Legacy.BluetoothNetwork.CancelPairingMode"
}
func EventNameType_LegacyBluetoothNetworkDeviceConnectedFailure() EventNameType {
	return "Legacy.BluetoothNetwork.DeviceConnectedFailure"
}
func EventNameType_LegacyBluetoothNetworkDeviceConnectedSuccess() EventNameType {
	return "Legacy.BluetoothNetwork.DeviceConnectedSuccess"
}
func EventNameType_LegacyBluetoothNetworkDeviceDisconnectedFailure() EventNameType {
	return "Legacy.BluetoothNetwork.DeviceDisconnectedFailure"
}
func EventNameType_LegacyBluetoothNetworkDeviceDisconnectedSuccess() EventNameType {
	return "Legacy.BluetoothNetwork.DeviceDisconnectedSuccess"
}
func EventNameType_LegacyBluetoothNetworkDevicePairFailure() EventNameType {
	return "Legacy.BluetoothNetwork.DevicePairFailure"
}
func EventNameType_LegacyBluetoothNetworkDevicePairSuccess() EventNameType {
	return "Legacy.BluetoothNetwork.DevicePairSuccess"
}
func EventNameType_LegacyBluetoothNetworkDeviceUnpairFailure() EventNameType {
	return "Legacy.BluetoothNetwork.DeviceUnpairFailure"
}
func EventNameType_LegacyBluetoothNetworkDeviceUnpairSuccess() EventNameType {
	return "Legacy.BluetoothNetwork.DeviceUnpairSuccess"
}
func EventNameType_LegacyBluetoothNetworkEnterPairingModeFailure() EventNameType {
	return "Legacy.BluetoothNetwork.EnterPairingModeFailure"
}
func EventNameType_LegacyBluetoothNetworkEnterPairingModeSuccess() EventNameType {
	return "Legacy.BluetoothNetwork.EnterPairingModeSuccess"
}
func EventNameType_LegacyBluetoothNetworkMediaControlFailure() EventNameType {
	return "Legacy.BluetoothNetwork.MediaControlFailure"
}
func EventNameType_LegacyBluetoothNetworkMediaControlSuccess() EventNameType {
	return "Legacy.BluetoothNetwork.MediaControlSuccess"
}
func EventNameType_LegacyBluetoothNetworkScanDevicesReport() EventNameType {
	return "Legacy.BluetoothNetwork.ScanDevicesReport"
}
func EventNameType_LegacyBluetoothNetworkSetDeviceCategoriesFailed() EventNameType {
	return "Legacy.BluetoothNetwork.SetDeviceCategoriesFailed"
}
func EventNameType_LegacyBluetoothNetworkSetDeviceCategoriesSucceeded() EventNameType {
	return "Legacy.BluetoothNetwork.SetDeviceCategoriesSucceeded"
}
func EventNameType_LegacyContentManagerContentPlaybackTerminated() EventNameType {
	return "Legacy.ContentManager.ContentPlaybackTerminated"
}
func EventNameType_LegacyDeviceNotificationDeleteNotificationFailed() EventNameType {
	return "Legacy.DeviceNotification.DeleteNotificationFailed"
}
func EventNameType_LegacyDeviceNotificationDeleteNotificationSucceeded() EventNameType {
	return "Legacy.DeviceNotification.DeleteNotificationSucceeded"
}
func EventNameType_LegacyDeviceNotificationNotificationEnteredBackground() EventNameType {
	return "Legacy.DeviceNotification.NotificationEnteredBackground"
}
func EventNameType_LegacyDeviceNotificationNotificationEnteredForground() EventNameType {
	return "Legacy.DeviceNotification.NotificationEnteredForground"
}
func EventNameType_LegacyDeviceNotificationNotificationStarted() EventNameType {
	return "Legacy.DeviceNotification.NotificationStarted"
}
func EventNameType_LegacyDeviceNotificationNotificationStopped() EventNameType {
	return "Legacy.DeviceNotification.NotificationStopped"
}
func EventNameType_LegacyDeviceNotificationNotificationSync() EventNameType {
	return "Legacy.DeviceNotification.NotificationSync"
}
func EventNameType_LegacyDeviceNotificationSetNotificationFailed() EventNameType {
	return "Legacy.DeviceNotification.SetNotificationFailed"
}
func EventNameType_LegacyDeviceNotificationSetNotificationSucceeded() EventNameType {
	return "Legacy.DeviceNotification.SetNotificationSucceeded"
}
func EventNameType_LegacyEqualizerControllerEqualizerChanged() EventNameType {
	return "Legacy.EqualizerController.EqualizerChanged"
}
func EventNameType_LegacyExternalMediaPlayerAuthorizationComplete() EventNameType {
	return "Legacy.ExternalMediaPlayer.AuthorizationComplete"
}
func EventNameType_LegacyExternalMediaPlayerError() EventNameType {
	return "Legacy.ExternalMediaPlayer.Error"
}
func EventNameType_LegacyExternalMediaPlayerEvent() EventNameType {
	return "Legacy.ExternalMediaPlayer.Event"
}
func EventNameType_LegacyExternalMediaPlayerLogin() EventNameType {
	return "Legacy.ExternalMediaPlayer.Login"
}
func EventNameType_LegacyExternalMediaPlayerLogout() EventNameType {
	return "Legacy.ExternalMediaPlayer.Logout"
}
func EventNameType_LegacyExternalMediaPlayerReportDiscoveredPlayers() EventNameType {
	return "Legacy.ExternalMediaPlayer.ReportDiscoveredPlayers"
}
func EventNameType_LegacyExternalMediaPlayerRequestToken() EventNameType {
	return "Legacy.ExternalMediaPlayer.RequestToken"
}
func EventNameType_LegacyFavoritesControllerError() EventNameType {
	return "Legacy.FavoritesController.Error"
}
func EventNameType_LegacyFavoritesControllerResponse() EventNameType {
	return "Legacy.FavoritesController.Response"
}
func EventNameType_LegacyGameEngineGameInputEvent() EventNameType {
	return "Legacy.GameEngine.GameInputEvent"
}
func EventNameType_LegacyHomeAutoWifiControllerDeviceReconnected() EventNameType {
	return "Legacy.HomeAutoWifiController.DeviceReconnected"
}
func EventNameType_LegacyHomeAutoWifiControllerHttpNotified() EventNameType {
	return "Legacy.HomeAutoWifiController.HttpNotified"
}
func EventNameType_LegacyHomeAutoWifiControllerSsdpDiscoveryFinished() EventNameType {
	return "Legacy.HomeAutoWifiController.SsdpDiscoveryFinished"
}
func EventNameType_LegacyHomeAutoWifiControllerSsdpServiceDiscovered() EventNameType {
	return "Legacy.HomeAutoWifiController.SsdpServiceDiscovered"
}
func EventNameType_LegacyHomeAutoWifiControllerSsdpServiceTerminated() EventNameType {
	return "Legacy.HomeAutoWifiController.SsdpServiceTerminated"
}
func EventNameType_LegacyListModelAddItemRequest() EventNameType {
	return "Legacy.ListModel.AddItemRequest"
}
func EventNameType_LegacyListModelDeleteItemRequest() EventNameType {
	return "Legacy.ListModel.DeleteItemRequest"
}
func EventNameType_LegacyListModelGetPageByOrdinalRequest() EventNameType {
	return "Legacy.ListModel.GetPageByOrdinalRequest"
}
func EventNameType_LegacyListModelGetPageByTokenRequest() EventNameType {
	return "Legacy.ListModel.GetPageByTokenRequest"
}
func EventNameType_LegacyListModelListStateUpdateRequest() EventNameType {
	return "Legacy.ListModel.ListStateUpdateRequest"
}
func EventNameType_LegacyListModelUpdateItemRequest() EventNameType {
	return "Legacy.ListModel.UpdateItemRequest"
}
func EventNameType_LegacyListRendererGetListPageByOrdinal() EventNameType {
	return "Legacy.ListRenderer.GetListPageByOrdinal"
}
func EventNameType_LegacyListRendererGetListPageByToken() EventNameType {
	return "Legacy.ListRenderer.GetListPageByToken"
}
func EventNameType_LegacyListRendererListItemEvent() EventNameType {
	return "Legacy.ListRenderer.ListItemEvent"
}
func EventNameType_LegacyMediaGroupingGroupChangeNotificationEvent() EventNameType {
	return "Legacy.MediaGrouping.GroupChangeNotificationEvent"
}
func EventNameType_LegacyMediaGroupingGroupChangeResponseEvent() EventNameType {
	return "Legacy.MediaGrouping.GroupChangeResponseEvent"
}
func EventNameType_LegacyMediaGroupingGroupSyncEvent() EventNameType {
	return "Legacy.MediaGrouping.GroupSyncEvent"
}
func EventNameType_LegacyMediaPlayerPlaybackError() EventNameType {
	return "Legacy.MediaPlayer.PlaybackError"
}
func EventNameType_LegacyMediaPlayerPlaybackFinished() EventNameType {
	return "Legacy.MediaPlayer.PlaybackFinished"
}
func EventNameType_LegacyMediaPlayerPlaybackIdle() EventNameType {
	return "Legacy.MediaPlayer.PlaybackIdle"
}
func EventNameType_LegacyMediaPlayerPlaybackNearlyFinished() EventNameType {
	return "Legacy.MediaPlayer.PlaybackNearlyFinished"
}
func EventNameType_LegacyMediaPlayerPlaybackPaused() EventNameType {
	return "Legacy.MediaPlayer.PlaybackPaused"
}
func EventNameType_LegacyMediaPlayerPlaybackResumed() EventNameType {
	return "Legacy.MediaPlayer.PlaybackResumed"
}
func EventNameType_LegacyMediaPlayerPlaybackStarted() EventNameType {
	return "Legacy.MediaPlayer.PlaybackStarted"
}
func EventNameType_LegacyMediaPlayerPlaybackStopped() EventNameType {
	return "Legacy.MediaPlayer.PlaybackStopped"
}
func EventNameType_LegacyMediaPlayerSequenceItemsRequested() EventNameType {
	return "Legacy.MediaPlayer.SequenceItemsRequested"
}
func EventNameType_LegacyMediaPlayerSequenceModified() EventNameType {
	return "Legacy.MediaPlayer.SequenceModified"
}
func EventNameType_LegacyMeetingClientControllerEvent() EventNameType {
	return "Legacy.MeetingClientController.Event"
}
func EventNameType_LegacyMicrophoneAudioRecording() EventNameType {
	return "Legacy.Microphone.AudioRecording"
}
func EventNameType_LegacyPhoneCallControllerEvent() EventNameType {
	return "Legacy.PhoneCallController.Event"
}
func EventNameType_LegacyPlaybackControllerButtonCommand() EventNameType {
	return "Legacy.PlaybackController.ButtonCommand"
}
func EventNameType_LegacyPlaybackControllerLyricsViewedEvent() EventNameType {
	return "Legacy.PlaybackController.LyricsViewedEvent"
}
func EventNameType_LegacyPlaybackControllerNextCommand() EventNameType {
	return "Legacy.PlaybackController.NextCommand"
}
func EventNameType_LegacyPlaybackControllerPauseCommand() EventNameType {
	return "Legacy.PlaybackController.PauseCommand"
}
func EventNameType_LegacyPlaybackControllerPlayCommand() EventNameType {
	return "Legacy.PlaybackController.PlayCommand"
}
func EventNameType_LegacyPlaybackControllerPreviousCommand() EventNameType {
	return "Legacy.PlaybackController.PreviousCommand"
}
func EventNameType_LegacyPlaybackControllerToggleCommand() EventNameType {
	return "Legacy.PlaybackController.ToggleCommand"
}
func EventNameType_LegacyPlaylistControllerErrorResponse() EventNameType {
	return "Legacy.PlaylistController.ErrorResponse"
}
func EventNameType_LegacyPlaylistControllerResponse() EventNameType {
	return "Legacy.PlaylistController.Response"
}
func EventNameType_LegacyPresentationPresentationDismissedEvent() EventNameType {
	return "Legacy.Presentation.PresentationDismissedEvent"
}
func EventNameType_LegacyPresentationPresentationUserEvent() EventNameType {
	return "Legacy.Presentation.PresentationUserEvent"
}
func EventNameType_LegacySconeRemoteControlNext() EventNameType {
	return "Legacy.SconeRemoteControl.Next"
}
func EventNameType_LegacySconeRemoteControlPlayPause() EventNameType {
	return "Legacy.SconeRemoteControl.PlayPause"
}
func EventNameType_LegacySconeRemoteControlPrevious() EventNameType {
	return "Legacy.SconeRemoteControl.Previous"
}
func EventNameType_LegacySconeRemoteControlVolumeDown() EventNameType {
	return "Legacy.SconeRemoteControl.VolumeDown"
}
func EventNameType_LegacySconeRemoteControlVolumeUp() EventNameType {
	return "Legacy.SconeRemoteControl.VolumeUp"
}
func EventNameType_LegacySipClientEvent() EventNameType {
	return "Legacy.SipClient.Event"
}
func EventNameType_LegacySoftwareUpdateCheckSoftwareUpdateReport() EventNameType {
	return "Legacy.SoftwareUpdate.CheckSoftwareUpdateReport"
}
func EventNameType_LegacySoftwareUpdateInitiateSoftwareUpdateReport() EventNameType {
	return "Legacy.SoftwareUpdate.InitiateSoftwareUpdateReport"
}
func EventNameType_LegacySpeakerMuteChanged() EventNameType {
	return "Legacy.Speaker.MuteChanged"
}
func EventNameType_LegacySpeakerVolumeChanged() EventNameType {
	return "Legacy.Speaker.VolumeChanged"
}
func EventNameType_LegacySpeechRecognizerWakeWordChanged() EventNameType {
	return "Legacy.SpeechRecognizer.WakeWordChanged"
}
func EventNameType_LegacySpeechSynthesizerSpeechFinished() EventNameType {
	return "Legacy.SpeechSynthesizer.SpeechFinished"
}
func EventNameType_LegacySpeechSynthesizerSpeechInterrupted() EventNameType {
	return "Legacy.SpeechSynthesizer.SpeechInterrupted"
}
func EventNameType_LegacySpeechSynthesizerSpeechStarted() EventNameType {
	return "Legacy.SpeechSynthesizer.SpeechStarted"
}
func EventNameType_LegacySpeechSynthesizerSpeechSynthesizerError() EventNameType {
	return "Legacy.SpeechSynthesizer.SpeechSynthesizerError"
}
func EventNameType_LegacySpotifyEvent() EventNameType {
	return "Legacy.Spotify.Event"
}
func EventNameType_LegacySystemUserInactivity() EventNameType {
	return "Legacy.System.UserInactivity"
}
func EventNameType_LegacyUDPControllerBroadcastResponse() EventNameType {
	return "Legacy.UDPController.BroadcastResponse"
}
func EventNameType_LocalApplicationAlexaTranslationLiveTranslationEvent() EventNameType {
	return "LocalApplication.Alexa.Translation.LiveTranslation.Event"
}
func EventNameType_LocalApplicationAlexaNotificationsEvent() EventNameType {
	return "LocalApplication.AlexaNotifications.Event"
}
func EventNameType_LocalApplicationAlexaPlatformTestSpeechletEvent() EventNameType {
	return "LocalApplication.AlexaPlatformTestSpeechlet.Event"
}
func EventNameType_LocalApplicationAlexaVisionEvent() EventNameType {
	return "LocalApplication.AlexaVision.Event"
}
func EventNameType_LocalApplicationAlexaVoiceLayerEvent() EventNameType {
	return "LocalApplication.AlexaVoiceLayer.Event"
}
func EventNameType_LocalApplicationAvaPhysicalShoppingEvent() EventNameType {
	return "LocalApplication.AvaPhysicalShopping.Event"
}
func EventNameType_LocalApplicationCalendarEvent() EventNameType {
	return "LocalApplication.Calendar.Event"
}
func EventNameType_LocalApplicationClosetEvent() EventNameType {
	return "LocalApplication.Closet.Event"
}
func EventNameType_LocalApplicationCommunicationsEvent() EventNameType {
	return "LocalApplication.Communications.Event"
}
func EventNameType_LocalApplicationDeviceMessagingEvent() EventNameType {
	return "LocalApplication.DeviceMessaging.Event"
}
func EventNameType_LocalApplicationDigitalDashEvent() EventNameType {
	return "LocalApplication.DigitalDash.Event"
}
func EventNameType_LocalApplicationFireflyShoppingEvent() EventNameType {
	return "LocalApplication.FireflyShopping.Event"
}
func EventNameType_LocalApplicationGalleryEvent() EventNameType {
	return "LocalApplication.Gallery.Event"
}
func EventNameType_LocalApplicationHHOPhotosEvent() EventNameType {
	return "LocalApplication.HHOPhotos.Event"
}
func EventNameType_LocalApplicationHomeAutomationMediaEvent() EventNameType {
	return "LocalApplication.HomeAutomationMedia.Event"
}
func EventNameType_LocalApplicationKnightContactsEvent() EventNameType {
	return "LocalApplication.KnightContacts.Event"
}
func EventNameType_LocalApplicationKnightHomeEvent() EventNameType {
	return "LocalApplication.KnightHome.Event"
}
func EventNameType_LocalApplicationKnightHomeThingsToTryEvent() EventNameType {
	return "LocalApplication.KnightHomeThingsToTry.Event"
}
func EventNameType_LocalApplicationLocalMediaPlayerEvent() EventNameType {
	return "LocalApplication.LocalMediaPlayer.Event"
}
func EventNameType_LocalApplicationLocalVoiceUIEvent() EventNameType {
	return "LocalApplication.LocalVoiceUI.Event"
}
func EventNameType_LocalApplicationMShopEvent() EventNameType {
	return "LocalApplication.MShop.Event"
}
func EventNameType_LocalApplicationMShopPurchasingEvent() EventNameType {
	return "LocalApplication.MShopPurchasing.Event"
}
func EventNameType_LocalApplicationNotificationsAppEvent() EventNameType {
	return "LocalApplication.NotificationsApp.Event"
}
func EventNameType_LocalApplicationPhotosEvent() EventNameType {
	return "LocalApplication.Photos.Event"
}
func EventNameType_LocalApplicationSentryEvent() EventNameType {
	return "LocalApplication.Sentry.Event"
}
func EventNameType_LocalApplicationSipClientEvent() EventNameType {
	return "LocalApplication.SipClient.Event"
}
func EventNameType_LocalApplicationSipUserAgentEvent() EventNameType {
	return "LocalApplication.SipUserAgent.Event"
}
func EventNameType_LocalApplicationtodoRendererEvent() EventNameType {
	return "LocalApplication.todoRenderer.Event"
}
func EventNameType_LocalApplicationVideoExperienceServiceEvent() EventNameType {
	return "LocalApplication.VideoExperienceService.Event"
}
func EventNameType_LocalApplicationWebVideoPlayerEvent() EventNameType {
	return "LocalApplication.WebVideoPlayer.Event"
}
func EventNameType_AlexaCameraPhotoCaptureControllerCancelCaptureFailed() EventNameType {
	return "Alexa.Camera.PhotoCaptureController.CancelCaptureFailed"
}
func EventNameType_AlexaCameraPhotoCaptureControllerCancelCaptureFinished() EventNameType {
	return "Alexa.Camera.PhotoCaptureController.CancelCaptureFinished"
}
func EventNameType_AlexaCameraPhotoCaptureControllerCaptureFailed() EventNameType {
	return "Alexa.Camera.PhotoCaptureController.CaptureFailed"
}
func EventNameType_AlexaCameraPhotoCaptureControllerCaptureFinished() EventNameType {
	return "Alexa.Camera.PhotoCaptureController.CaptureFinished"
}
func EventNameType_AlexaCameraVideoCaptureControllerCancelCaptureFailed() EventNameType {
	return "Alexa.Camera.VideoCaptureController.CancelCaptureFailed"
}
func EventNameType_AlexaCameraVideoCaptureControllerCancelCaptureFinished() EventNameType {
	return "Alexa.Camera.VideoCaptureController.CancelCaptureFinished"
}
func EventNameType_AlexaCameraVideoCaptureControllerCaptureFailed() EventNameType {
	return "Alexa.Camera.VideoCaptureController.CaptureFailed"
}
func EventNameType_AlexaCameraVideoCaptureControllerCaptureFinished() EventNameType {
	return "Alexa.Camera.VideoCaptureController.CaptureFinished"
}
func EventNameType_AlexaCameraVideoCaptureControllerCaptureStarted() EventNameType {
	return "Alexa.Camera.VideoCaptureController.CaptureStarted"
}
func EventNameType_AlexaFileManagerUploadControllerCancelUploadFailed() EventNameType {
	return "Alexa.FileManager.UploadController.CancelUploadFailed"
}
func EventNameType_AlexaFileManagerUploadControllerCancelUploadFinished() EventNameType {
	return "Alexa.FileManager.UploadController.CancelUploadFinished"
}
func EventNameType_AlexaFileManagerUploadControllerUploadFailed() EventNameType {
	return "Alexa.FileManager.UploadController.UploadFailed"
}
func EventNameType_AlexaFileManagerUploadControllerUploadFinished() EventNameType {
	return "Alexa.FileManager.UploadController.UploadFinished"
}
func EventNameType_AlexaFileManagerUploadControllerUploadStarted() EventNameType {
	return "Alexa.FileManager.UploadController.UploadStarted"
}
func EventNameType_AlexaPresentationAPLLoadIndexListData() EventNameType {
	return "Alexa.Presentation.APL.LoadIndexListData"
}
func EventNameType_AlexaPresentationAPLRuntimeError() EventNameType {
	return "Alexa.Presentation.APL.RuntimeError"
}
func EventNameType_AlexaPresentationAPLUserEvent() EventNameType {
	return "Alexa.Presentation.APL.UserEvent"
}
func EventNameType_AlexaPresentationHTMLEvent() EventNameType {
	return "Alexa.Presentation.HTML.Event"
}
func EventNameType_AlexaPresentationHTMLLifecycleStateChanged() EventNameType {
	return "Alexa.Presentation.HTML.LifecycleStateChanged"
}
func EventNameType_AlexaPresentationPresentationDismissed() EventNameType {
	return "Alexa.Presentation.PresentationDismissed"
}
func EventNameType_AudioPlayerPlaybackFailed() EventNameType {
	return "AudioPlayer.PlaybackFailed"
}
func EventNameType_AudioPlayerPlaybackFinished() EventNameType {
	return "AudioPlayer.PlaybackFinished"
}
func EventNameType_AudioPlayerPlaybackNearlyFinished() EventNameType {
	return "AudioPlayer.PlaybackNearlyFinished"
}
func EventNameType_AudioPlayerPlaybackStarted() EventNameType {
	return "AudioPlayer.PlaybackStarted"
}
func EventNameType_AudioPlayerPlaybackStopped() EventNameType {
	return "AudioPlayer.PlaybackStopped"
}
func EventNameType_CardRendererDisplayContentFinished() EventNameType {
	return "CardRenderer.DisplayContentFinished"
}
func EventNameType_CardRendererDisplayContentStarted() EventNameType {
	return "CardRenderer.DisplayContentStarted"
}
func EventNameType_CardRendererReadContentFinished() EventNameType {
	return "CardRenderer.ReadContentFinished"
}
func EventNameType_CardRendererReadContentStarted() EventNameType {
	return "CardRenderer.ReadContentStarted"
}
func EventNameType_CustomInterfaceControllerEventsReceived() EventNameType {
	return "CustomInterfaceController.EventsReceived"
}
func EventNameType_CustomInterfaceControllerExpired() EventNameType {
	return "CustomInterfaceController.Expired"
}
func EventNameType_DeviceSetupSetupCompleted() EventNameType {
	return "DeviceSetup.SetupCompleted"
}
func EventNameType_DisplayElementSelected() EventNameType {
	return "Display.ElementSelected"
}
func EventNameType_DisplayUserEvent() EventNameType {
	return "Display.UserEvent"
}
func EventNameType_FitnessSessionControllerFitnessSessionEnded() EventNameType {
	return "FitnessSessionController.FitnessSessionEnded"
}
func EventNameType_FitnessSessionControllerFitnessSessionError() EventNameType {
	return "FitnessSessionController.FitnessSessionError"
}
func EventNameType_FitnessSessionControllerFitnessSessionPaused() EventNameType {
	return "FitnessSessionController.FitnessSessionPaused"
}
func EventNameType_FitnessSessionControllerFitnessSessionResumed() EventNameType {
	return "FitnessSessionController.FitnessSessionResumed"
}
func EventNameType_FitnessSessionControllerFitnessSessionStarted() EventNameType {
	return "FitnessSessionController.FitnessSessionStarted"
}
func EventNameType_GameEngineInputHandlerEvent() EventNameType {
	return "GameEngine.InputHandlerEvent"
}
func EventNameType_MessagingMessageReceived() EventNameType {
	return "Messaging.MessageReceived"
}
func EventNameType_MessagingControllerUpdateConversationsStatus() EventNameType {
	return "MessagingController.UpdateConversationsStatus"
}
func EventNameType_MessagingControllerUpdateMessagesStatusRequest() EventNameType {
	return "MessagingController.UpdateMessagesStatusRequest"
}
func EventNameType_MessagingControllerUpdateSendMessageStatusRequest() EventNameType {
	return "MessagingController.UpdateSendMessageStatusRequest"
}
func EventNameType_MessagingControllerUploadConversations() EventNameType {
	return "MessagingController.UploadConversations"
}
func EventNameType_PlaybackControllerNextCommandIssued() EventNameType {
	return "PlaybackController.NextCommandIssued"
}
func EventNameType_PlaybackControllerPauseCommandIssued() EventNameType {
	return "PlaybackController.PauseCommandIssued"
}
func EventNameType_PlaybackControllerPlayCommandIssued() EventNameType {
	return "PlaybackController.PlayCommandIssued"
}
func EventNameType_PlaybackControllerPreviousCommandIssued() EventNameType {
	return "PlaybackController.PreviousCommandIssued"
}
func EventNameType_EffectsControllerRequestEffectChangeRequest() EventNameType {
	return "EffectsController.RequestEffectChangeRequest"
}
func EventNameType_EffectsControllerRequestGuiChangeRequest() EventNameType {
	return "EffectsController.RequestGuiChangeRequest"
}
func EventNameType_EffectsControllerStateReceiptChangeRequest() EventNameType {
	return "EffectsController.StateReceiptChangeRequest"
}
func EventNameType_AlexaVideoXrayShowDetailsSuccessful() EventNameType {
	return "Alexa.Video.Xray.ShowDetailsSuccessful"
}
func EventNameType_AlexaVideoXrayShowDetailsFailed() EventNameType {
	return "Alexa.Video.Xray.ShowDetailsFailed"
}
