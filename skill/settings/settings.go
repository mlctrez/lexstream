package settings

import (
	"fmt"
	"github.com/mlctrez/lexstream/internal/jutil"
	"github.com/mlctrez/lexstream/smapi"
)

type Settings struct {
	Bucket     string `json:"bucket"`
	LambdaName string `json:"lambda_name"`
	SkillName  string `json:"skill_name"`
	SkillId    string `json:"-"`
}

func Load() (s *Settings, err error) {
	s = &Settings{}

	if err = jutil.DecodePath("settings.json", s); err != nil {
		return
	}

	if s.Bucket == "" {
		return nil, fmt.Errorf("no bucket name in settings.json")
	}

	if s.SkillId == "" {
		id, getSkillIdErr := smapi.New().GetSkillIdForName(s.SkillName)
		if getSkillIdErr != nil {
			return nil, getSkillIdErr
		}
		s.SkillId = id
	}

	return
}
