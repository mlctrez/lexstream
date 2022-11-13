package smapi

import (
	"context"
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
