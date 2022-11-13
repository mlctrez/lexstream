package types

// EntityType represents EntityTypes supported by the Music Kill API.
type EntityType string

// TRACK – a song like Lady Gaga's "Gypsy"
const TRACK EntityType = "TRACK"

// ALBUM – an album like Lady Gaga's "ARTPOP"
const ALBUM EntityType = "ALBUM"

// ARTIST – an artist like "Lady Gaga"
const ARTIST EntityType = "ARTIST"

// PLAYLIST – a playlist like "Relaxing Sounds"
const PLAYLIST EntityType = "PLAYLIST"

// GENRE – a genre like "Jazz"
const GENRE EntityType = "GENRE"

// STATION – a station like "C.N.N."
const STATION EntityType = "STATION"

// MEDIA_TYPE – describes different types of audio content, for example, SONG or PROGRAM_SERIES
const MEDIA_TYPE EntityType = "MEDIA_TYPE"

// RELEASE_WINDOW – (Podcast only) the release date window
const RELEASE_WINDOW EntityType = "RELEASE_WINDOW"

// SORT_TYPE – describes whether audio contents or selection criteria should be ranked by
// POPULARITY or RECENCY before being returned to Alexa
const SORT_TYPE EntityType = "SORT_TYPE"

// FOLLOW_STATUS – (Podcast only) indicates whether the user follows or subscribes to a program series
const FOLLOW_STATUS EntityType = "FOLLOW_STATUS"

// TOPIC_NAME - (Podcast only) describes the topic of a program series. See Implement episode search by topic.
const TOPIC_NAME EntityType = "TOPIC_NAME"
