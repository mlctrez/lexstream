package smapi

import (
	"context"
	"fmt"
	v0 "github.com/mlctrez/lexstream/smapiv0/client"
	v1 "github.com/mlctrez/lexstream/smapiv1/client"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"time"
)

const Endpoint = "https://api.amazonalexa.com"

func New() *Smapi {

	client := oauth2.NewClient(context.TODO(), &TokenSource{})

	client.Transport = &DelayRoundTripper{
		Target:      client.Transport,
		Delay:       time.Second * 1,
		LogRequests: false,
	}

	return &Smapi{
		V0: &v0.Client{Client: client, Endpoint: Endpoint},
		V1: &v1.Client{Client: client, Endpoint: Endpoint},
	}
}

var _ http.RoundTripper = (*DelayRoundTripper)(nil)

type DelayRoundTripper struct {
	Target      http.RoundTripper
	Delay       time.Duration
	LogRequests bool
}

func (d *DelayRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	time.Sleep(d.Delay)
	if d.LogRequests {
		log.Println("request", request.Method, request.URL.Path)
	}
	response, err := d.Target.RoundTrip(request)
	if d.LogRequests {
		log.Println("response", response.Status)
	}
	return response, err
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

func (s *Smapi) GetSingleVendor() (string, error) {
	response, err := s.V1.GetVendorListV1()
	if err != nil {
		return "", err
	}
	var vendorId string
	for _, vendor := range response.Vendors {
		if vendorId != "" {
			return "", fmt.Errorf("more than one vendor found")
		}
		vendorId = vendor.Id
	}
	if vendorId == "" {
		return "", fmt.Errorf("no vendors found")
	}
	return vendorId, nil

}
