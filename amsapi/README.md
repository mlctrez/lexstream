# Namespace / Operation to file reference

Completed models have links

## Audio.PlayQueue

### Directives

* GetNextItem directive
* GetPreviousItem directive
* JumpToItem directive

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
