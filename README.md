# lexstream

Alexa Music skill for streaming your own music.

[![Go Report Card](https://goreportcard.com/badge/github.com/mlctrez/lexstream)](https://goreportcard.com/report/github.com/mlctrez/lexstream)

## Documentation links

Various documentation links that have been useful when building out this project

* [how-a-music-skill-works](https://developer.amazon.com/en-US/docs/alexa/music-skills/understand-the-music-skill-api.html#how-a-music-skill-works)

### Prerequisites

* [Go](https://go.dev/doc/install)
* [Node.js](https://nodejs.org/en/download/)
  * Alexa Skills Kit requires `node --version` v8.3 or higher.
* AWS Account
  * [Configuration and credential file settings](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)
  * These credentials need to have permissions to change S3, IAM, DynamoDB, and Lambda resources
* [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
* [Alexa Developer Console](https://developer.amazon.com/alexa/console/ask)
  * Validate that you can sign in

### Initialize Go Modules

* Use `go mod tidy` in the root of this repository to initialize dependencies

### Run preflight check

* `go run cmd/preflight.go` will:
  * Check that the minimum version of node is installed.
  * Configure and test ask-cli in the .ask directory.
    * This will prompt to run `ask configure` if no credentials are found 
  * Check that AWS credentials and configuration files are set up correctly.
    * This performs read only api calls for S3, IAM, DynamoDB, and Lambda 
  * Check that the aws cli command is installed

> **Warning** commands after this point start modifying AWS resources, read closely to warnings

### Deploy Skill

TODO

### Deploy Lambda

* WARNING the bucket permissions will be re-set to no public access!!
* Copy settings.json.example to settings.json and update with the appropriate content
* Run the following command to create the required Bucket, IAM Role, and Lambda
    ```shell
  $ go run infra/sync.go
    ```
### Update Endpoint and Re-Validate

TODO: this should be able to be automated
* Back in the Alexa skills console, update the skill endpoint with the output of the previous command

### Upload Design Notes

* Use [MusicBrainz_API](https://musicbrainz.org/doc/MusicBrainz_API) for identifiers wherever possible.
  * Completed
  * Artist,Album, and Track are all directly supported and would be unique
  * Prefix each since the MB Api uses GUIDs that are opaque

### TODOs and other ramblings

* Right now this just plays a single song on loop
* Create a web interface / cli to upload new songs
    * upload audio files and cover art to s3
    * extract metadata from id tags in audio and create catalog files
    * automate upload of catalog data
        * ~~The `ask cli` does a browser login to get the access token so this may not be possible~~
        * [more research](https://developer.amazon.com/en-US/docs/alexa/smapi/get-access-token-smapi.html)
        * This is possible using LWA tokens and calling the upload catalog api directly
        * Add documentation on how this works

* Metadata notes showing id -> api mappings
    * [MUSICBRAINZ_ARTISTID](https://musicbrainz.org/ws/2/artist/7944ed53-2a58-4035-9b93-140a71e41c34?fmt=json)
    * [MUSICBRAINZ_ALBUMARTISTID](https://musicbrainz.org/ws/2/artist/7944ed53-2a58-4035-9b93-140a71e41c34?fmt=json)
    * [MUSICBRAINZ_TRACKID](https://musicbrainz.org/ws/2/recording/de26e48e-1b04-46ad-aa32-8a19a038c173?fmt=json)
    * [MUSICBRAINZ_ALBUMID](https://musicbrainz.org/ws/2/release/fe425df3-1844-397d-95b3-a85528aa98d7?fmt=json&inc=recordings+release-groups)
        * This ID can be used on
          the [cover art archive](https://coverartarchive.org/release/fe425df3-1844-397d-95b3-a85528aa98d7)
        * Redirects to
          the [internet archive](https://ia802607.us.archive.org/26/items/mbid-fe425df3-1844-397d-95b3-a85528aa98d7/index.json)
    * [MUSICBRAINZ_RELEASEGROUPID](https://musicbrainz.org/ws/2/release-group/93ef0ae1-0735-3699-b450-d79bdcb3d0b8?fmt=json)
    * [MUSICBRAINZ_WORKID](https://musicbrainz.org/ws/2/work/ba7d0c33-cb5b-3146-8a51-9d9de8f17ad3?fmt=json)
    * MUSICBRAINZ_RELEASETRACKID=a022f133-009d-3e9f-b0c1-d2e30e86e98e
        * I don't think this one can be directly queried, but it shows up under /media/tracks/id in
          [MUSICBRAINZ_ALBUMID](https://musicbrainz.org/ws/2/release/fe425df3-1844-397d-95b3-a85528aa98d7?fmt=json&inc=recordings+release-groups)
        * It can also be used as an anchor for
          the [release page](https://musicbrainz.org/release/fe425df3-1844-397d-95b3-a85528aa98d7/disc/1#a022f133-009d-3e9f-b0c1-d2e30e86e98e)

* Saving whipper command `whipper cd rip -p`

Initial readme created by [tigwen](https://github.com/mlctrez/tigwen)
