package main

import (
	"fmt"
	"github.com/mlctrez/lexstream/skill/settings"
	"github.com/mlctrez/lexstream/smapi"
	"github.com/mlctrez/lexstream/smapiv1/skill"
	manifest "github.com/mlctrez/lexstream/smapiv1/skill/Manifest"
	"log"
	"strings"
	"time"
)

type CreateSkill struct {
	smapi    *smapi.Smapi
	settings *settings.Settings
}

func logAndExit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (cs *CreateSkill) VendorID() (vendorID string) {
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

func (cs *CreateSkill) getSkills() []*skill.SkillSummary {
	resp, err := cs.smapi.V1.ListSkillsForVendorV1(cs.VendorID(), "", 50, []string{})
	if err != nil {
		log.Fatal(err)
	}
	return resp.Skills

}

func (cs *CreateSkill) Run() (err error) {

	if cs.settings, err = settings.Load(); err != nil {
		return
	}

	for _, summary := range cs.getSkills() {
		name := summary.NameByLocale["en-US"]
		if name == cs.settings.SkillName {
			fmt.Println("deleting existing skill with id", summary.SkillId)
			if err := cs.smapi.V1.DeleteSkillV1(summary.SkillId); err != nil {
				log.Fatal(err)
			}
		}
	}

	response := cs.createSkill(cs.settings.SkillName)
	skillFound := false
	for !skillFound {
		for _, summary := range cs.getSkills() {
			if summary.SkillId == response.SkillId {
				skillFound = true
			}
		}
		if !skillFound {
			fmt.Println("waiting for skill to appear in ListSkillsForVendorV1")
			time.Sleep(3 * time.Second)
		}
	}

	for _, summary := range cs.getSkills() {
		name := summary.NameByLocale["en-US"]
		fmt.Println(summary.SkillId, name)
	}
	return nil
}

func (cs *CreateSkill) createSkill(skillName string) *skill.CreateSkillResponse {
	distMode := manifest.DistributionMode_PUBLIC()
	response, err := cs.smapi.V1.CreateSkillForVendorV1(&skill.CreateSkillRequest{
		Manifest: &manifest.SkillManifest{
			ManifestVersion: "1.0",
			PublishingInformation: &manifest.SkillManifestPublishingInformation{
				Locales: map[string]manifest.SkillManifestLocalizedPublishingInformation{
					"en-US": {Name: skillName},
				},
				IsAvailableWorldwide: true,
				DistributionMode:     &distMode,
			},
			Apis: &manifest.SkillManifestApis{
				Music: &manifest.MusicApis{
					Interfaces: MusicInterfaces(),
					Locales: map[string]manifest.LocalizedMusicInfo{
						"en-US": {
							PromptName: skillName,
							Aliases:    []*manifest.MusicAlias{{Name: strings.ToLower(skillName)}},
							Features:   []*manifest.MusicFeature{{Name: "EXPLICIT_LANGUAGE_FILTER"}},
						},
					},
				},
			},
		},
		VendorId: cs.VendorID(),
	})
	if err != nil {
		log.Fatal(err)
	}
	return response
}

func MusicInterfaces() []*manifest.MusicInterfaces {

	return []*manifest.MusicInterfaces{
		{
			Namespace: "Alexa.Media.Playback",
			Version:   "1.0",
			Requests:  []*manifest.MusicRequest{{Name: "Initiate"}},
		},
		{
			Namespace: "Alexa.Media.Search",
			Version:   "1.0",
			Requests:  []*manifest.MusicRequest{{Name: "GetPlayableContent"}},
		},
		{
			Namespace: "Alexa.Audio.PlayQueue",
			Version:   "1.0",
			Requests: []*manifest.MusicRequest{
				{Name: "GetNextItem"},
				{Name: "GetPreviousItem"},
			},
		},
	}

}

func main() {

	cs := &CreateSkill{smapi: smapi.New()}
	err := cs.Run()
	if err != nil {
		log.Fatal(err)
	}

}
