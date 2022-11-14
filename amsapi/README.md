# Namespace / Operation to file reference

Completed models have links

## Audio.PlayQueue

### Directives

* GetNextItem directive [GetNextItem](playqueue.go#L3)
  * response [GetNextItemResponse](playqueue.go#L15)
* GetPreviousItem directive [GetPreviousItem](playqueue.go#L28)
  * response [GetPreviousItemResponse](playqueue.go#L37)
* JumpToItem directive [JumpToItem](playqueue.go#L44)
  * response [JumpToItemResponse](playqueue.go#L57)

## Media.Playback

### Enable playback directives

* Initiate directive [Initiate](playback.go#L5)
    * response [InitiateResponse](playback.go#L46)
* ReInitiate directive

## Media.PlayQueue

### Item Control directives

* GetItem 
* GetView 

### Playback directives

* Loop 
* Repeat 
* Shuffle 

## Media.Search

### Directives

* GetPlayableContent directive [GetPlayableContent](search.go#L5)
  * [GetPlayableContentResponse](search.go#L72)
