package smapi

import (
	"context"
	"fmt"
	v0 "github.com/mlctrez/lexstream/smapiv0/client"
	v1 "github.com/mlctrez/lexstream/smapiv1/client"
	"golang.org/x/oauth2"
)

const Endpoint = "https://api.amazonalexa.com"

func Version0() *v0.Client {
	return &v0.Client{Client: oauth2.NewClient(context.TODO(), &TokenSource{}), Endpoint: Endpoint}

}

func Version1() *v1.Client {
	return &v1.Client{Client: oauth2.NewClient(context.TODO(), &TokenSource{}), Endpoint: Endpoint}
}

type Smapi struct {
	V0 *v0.Client
	V1 *v1.Client
}

func (s *Smapi) GetSkillIdForName(name string) (id string, err error) {
	response, err := s.V1.GetVendorListV1()
	if err != nil {
		return "", err
	}
	if len(response.Vendors) != 1 {
		return "", fmt.Errorf("unable to get 1 and only 1 vendor")
	}

	listResponse, err := s.V1.ListSkillsForVendorV1(response.Vendors[0].Id, "", 0, []string{})
	if err != nil {
		return "", err
	}

	for _, skill := range listResponse.Skills {
		if skill.NameByLocale["en-US"] == name {
			return skill.SkillId, nil
		}
	}
	return "", fmt.Errorf("unable to find skill id for name %s", name)
}

func New() *Smapi {
	return &Smapi{V0: Version0(), V1: Version1()}
}
