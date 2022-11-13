package catalog

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type Builder struct {
	ArtistCatalog *ArtistCatalog
	AlbumCatalog  *AlbumCatalog
	TrackCatalog  *TrackCatalog
}

func New() *Builder {
	return &Builder{
		ArtistCatalog: CreateArtistCatalog(),
		AlbumCatalog:  CreateAlbumCatalog(),
		TrackCatalog:  CreateTrackCatalog(),
	}
}

func (b *Builder) Walk(dir string) error {
	return filepath.Walk(dir, b.walkFunc)
}

func (b *Builder) walkFunc(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	lastModified := info.ModTime()

	switch filepath.Ext(info.Name()) {
	case ".flac":
		var md MetaData
		if md, err = flacMetaDataFromPath(path); err != nil {
			return err
		}

		if md.AlbumID() == "" || md.ArtistID() == "" || md.TrackID() == "" {
			missingError := fmt.Errorf("missing metadata on %q", path)
			// TODO: ignore always, fail, or prompt for cleanup
			fmt.Println(missingError)
			return nil
		}

		b.ArtistCatalog.AddMetaData(md, lastModified)
		b.AlbumCatalog.AddMetaData(md, lastModified)
		b.TrackCatalog.AddMetaData(md, lastModified)

	case ".ogg", ".mp3":
		// TODO: make ogg and mp3 md scanners
		return fmt.Errorf("unsupported file %q", path)
	default:
		// ignore all other files
	}
	return nil
}

// MetaData represents the musicbrainz data extracted from VORBIS_COMMENT
type MetaData interface {
	// Artist is the tag data ARTIST
	Artist() string
	// ArtistID is the tag data MUSICBRAINZ_ARTISTID
	ArtistID() string
	// Album is the tag data ALBUM
	Album() string
	// AlbumID is the tag data MUSICBRAINZ_ALBUMID
	AlbumID() string
	// Track is the tag data TITLE
	Track() string
	// TrackID is the tag data MUSICBRAINZ_TRACKID
	TrackID() string
}

var _ MetaData = (*metaData)(nil)

type metaData map[string]string

func (m metaData) Artist() string   { return m["ARTIST"] }
func (m metaData) ArtistID() string { return m["MUSICBRAINZ_ARTISTID"] }
func (m metaData) Album() string    { return m["ALBUM"] }
func (m metaData) AlbumID() string  { return m["MUSICBRAINZ_ALBUMID"] }
func (m metaData) Track() string    { return m["TITLE"] }
func (m metaData) TrackID() string  { return m["MUSICBRAINZ_TRACKID"] }

func flacMetaDataFromPath(path string) (md MetaData, err error) {
	command := exec.Command("metaflac", "--list", "--block-type=VORBIS_COMMENT", path)
	var output []byte
	output, err = command.CombinedOutput()
	if err != nil {
		return
	}
	m := make(map[string]string)
	scanner := bufio.NewScanner(bytes.NewBuffer(output))
	reg := regexp.MustCompile(`^comment\[\d+]: (\w+)=(.*)`)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		//fmt.Println(line)
		matches := reg.FindStringSubmatch(line)
		if len(matches) == 3 {
			m[matches[1]] = matches[2]
			//fmt.Println(strings.Join(matches[1:], "++"))
		}
	}
	err = scanner.Err()
	if err != nil {
		md = metaData(map[string]string{})
	}
	md = metaData(m)
	return
}

type MetaDataReceiver interface {
	AddMetaData(metaData MetaData, lastUpdate time.Time)
}

var exampleFlacMetadata = `
{
  "ALBUM": "Lay It Down",
  "ALBUMARTIST": "Cowboy Junkies",
  "ARTIST": "Cowboy Junkies",
  "DATE": "1996-02-27",
  "DISCNUMBER": "1",
  "DISCTOTAL": "1",
  "ISRC": "USGF19595201",
  "MUSICBRAINZ_ALBUMARTISTID": "b52211f1-49a2-4cd2-8535-dd1ba37ea90b",
  "MUSICBRAINZ_ALBUMID": "e04e5f1c-e502-47b9-9f5c-04dc26ce922d",
  "MUSICBRAINZ_ARTISTID": "b52211f1-49a2-4cd2-8535-dd1ba37ea90b",
  "MUSICBRAINZ_DISCID": "4PT.4pO4phCEVCk1NeLRDz9JAJQ-",
  "MUSICBRAINZ_RELEASEGROUPID": "67fdf1a3-2127-3de6-8d91-dd4b8ac6abf0",
  "MUSICBRAINZ_RELEASETRACKID": "7cf7e624-4dd7-3595-bbb6-b692a03eb0ad",
  "MUSICBRAINZ_TRACKID": "aaa0d8db-e0ed-44f1-8e70-51514ab65056",
  "MUSICBRAINZ_WORKID": "1120afd0-7305-4cf1-a86a-12cebf850c3a",
  "TITLE": "Something More Besides You",
  "TRACKNUMBER": "1",
  "TRACKTOTAL": "13"
}`
