package main

import (
	"github.com/mlctrez/lexstream/internal/jutil"
	"github.com/mlctrez/lexstream/internal/settings"
	"github.com/mlctrez/lexstream/smapi"
	"github.com/mlctrez/lexstream/smapiv1/skill"
	"log"
	"os"
)

type GetManifest struct {
	smapi    *smapi.Smapi
	settings *settings.Settings
}

func logAndExit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (cs *GetManifest) VendorID() (vendorID string) {
	if response, err := cs.smapi.V1.GetVendorListV1(); err != nil {
		logAndExit(err)
		return ""
	} else {
		if len(response.Vendors) != 1 {
			// TODO: be able to configure the preferred vendor to use
			log.Fatal("exiting since exactly 1 vendor was not found")
		}
		return response.Vendors[0].Id
	}
}

func (cs *GetManifest) getSkills() []*skill.SkillSummary {
	resp, err := cs.smapi.V1.ListSkillsForVendorV1(cs.VendorID(), "", 50, []string{})
	if err != nil {
		log.Fatal(err)
	}
	return resp.Skills
}

func (cs *GetManifest) Run() (err error) {

	if cs.settings, err = settings.Load(); err != nil {
		return
	}

	for _, summary := range cs.getSkills() {
		name := summary.NameByLocale["en-US"]
		if name == cs.settings.SkillName {
			response, err := cs.smapi.V1.GetSkillManifestV1(summary.SkillId, "development")
			logAndExit(err)
			buf, err := jutil.Marshal(response.Manifest, true)
			logAndExit(err)
			logAndExit(os.MkdirAll("temp", 0755))
			logAndExit(os.WriteFile("temp/manifest.json", buf.Bytes(), 0644))
		}
	}

	return nil
}

func main() {

	cs := &GetManifest{smapi: smapi.New()}
	err := cs.Run()
	if err != nil {
		log.Fatal(err)
	}

}
