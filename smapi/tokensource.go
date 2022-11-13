package smapi

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/mlctrez/lexstream/internal/jutil"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/amazon"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

const TokenPath = "token.json"
const LwaClientData = "secret.json"

var _ oauth2.TokenSource = (*TokenSource)(nil)

type TokenSource struct{}

// Token provides the bearer token for api calls to the alexa api Endpoint
func (t *TokenSource) Token() (token *oauth2.Token, err error) {

	token = &oauth2.Token{}
	if err = jutil.DecodePath(TokenPath, token); err != nil {
		return nil, err
	}

	// token read from disk was still valid, use it instead
	if token.Valid() {
		return
	}

	secret := &ClientCredentials{}
	if err = jutil.DecodePath(LwaClientData, secret); err != nil {
		return nil, err
	}

	requestToken := &RefreshTokenRequest{
		GrantType:    "refresh_token",
		ClientId:     secret.ClientId,
		ClientSecret: secret.ClientSecret,
		RefreshToken: token.RefreshToken,
	}

	var buf *bytes.Buffer
	buf, err = jutil.Marshal(requestToken, false)

	var request *http.Request
	if request, err = http.NewRequest("POST", amazon.Endpoint.TokenURL, buf); err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	var response *http.Response
	if response, err = http.DefaultClient.Do(request); err != nil {
		return
	}

	if response.StatusCode < 200 && response.StatusCode > 299 {
		err = fmt.Errorf("bad response code %d : %s", response.StatusCode, response.Status)
		return
	}

	responseToken := &RefreshTokenResponse{}
	if err = jutil.DecodeReadCloser(response.Body, responseToken); err != nil {
		return
	}

	token = &oauth2.Token{
		AccessToken:  responseToken.AccessToken,
		TokenType:    responseToken.TokenType,
		RefreshToken: responseToken.RefreshToken,
		Expiry:       time.Now().UTC().Add(time.Duration(responseToken.ExpiresIn) * time.Second),
	}
	err = jutil.WriteJson(TokenPath, token, true)

	return

}

// GenerateLwaTokens calls the ask util generate-lwa-tokens and saves the token data to token.json
// where it can be used by the token source.
//
// Follow the setup instructions at the following url, stopping before "Generate an access token".
//
// https://developer.amazon.com/en-US/docs/alexa/smapi/get-access-token-smapi.html#configure-lwa-security-profile
//
//	Save the client_id and client_secret in secrets.json with the following format:
//	{
//	 "client_id": "ID",
//	 "client_secret": "SECRET"
//	}
func GenerateLwaTokens() error {

	lwaClientCredentials := &ClientCredentials{}
	err := jutil.DecodePath(LwaClientData, lwaClientCredentials)
	if err != nil {
		return err
	}

	args := []string{
		"util", "generate-lwa-tokens",
		"--client-id", lwaClientCredentials.ClientId,
		"--client-confirmation", lwaClientCredentials.ClientSecret,
	}

	cmd := exec.Command(".ask/node_modules/.bin/ask", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return err
	}

	fmt.Println(string(output))

	jsonOnly := &bytes.Buffer{}

	scanner := bufio.NewScanner(bytes.NewBuffer(output))
	inJson := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "{") {
			inJson = true
		}
		if inJson {
			jsonOnly.WriteString(line)
		}
	}
	err = scanner.Err()
	if err != nil {
		return err
	}
	res := &RefreshTokenResponse{}

	err = jutil.DecodeReader(jsonOnly, res)
	if err != nil {
		return err
	}

	return jutil.WriteJson(TokenPath, res, true)
}

type RefreshTokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int       `json:"expires_in"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type ClientCredentials struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
