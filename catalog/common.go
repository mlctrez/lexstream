package catalog

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Header struct {
	Type    string   `json:"type"`
	Version float64  `json:"version"`
	Locales []Locale `json:"locales"`
}

type HeaderType string

const MusicAlbum HeaderType = "AMAZON.MusicAlbum"
const MusicGroup HeaderType = "AMAZON.MusicGroup"
const MusicRecording HeaderType = "AMAZON.MusicRecording"

func buildHeader(headerType HeaderType) Header {
	return Header{Type: string(headerType), Version: 2.0, Locales: DefaultLocales()}
}

type Locale struct {
	Country  string `json:"country"`
	Language string `json:"language"`
}

func DefaultLocales() []Locale {
	return []Locale{{Country: "US", Language: "en"}}
}

type Name struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

type AlternateName struct {
	Language string   `json:"language"`
	Values   []string `json:"values"`
}

type Popularity struct {
	Default   float64              `json:"default"`
	Overrides []PopularityOverride `json:"overrides,omitempty"`
}

type PopularityOverride struct {
	Locale Locale  `json:"locale"`
	Value  float64 `json:"value"`
}

var _ json.Marshaler = (*JSONTime)(nil)
var _ json.Unmarshaler = (*JSONTime)(nil)

type JSONTime time.Time

const jsonTimeFormat = "2006-01-02T15:04:05.999Z"

func (jt *JSONTime) UnmarshalJSON(bytes []byte) (err error) {
	s := strings.Trim(string(bytes), `"`)
	var nt time.Time
	nt, err = time.Parse(jsonTimeFormat, s)
	*jt = JSONTime(nt)
	return
}

func (jt *JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(jt.String()), nil
}

func (jt *JSONTime) String() string {
	t := time.Time(*jt)
	return fmt.Sprintf("%q", t.Format(jsonTimeFormat))
}
