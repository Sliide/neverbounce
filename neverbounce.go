package neverbounce

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const DEFAULT_API_URL string = "https://api.neverbounce.com/v3"

const (
	EMAIL_VALID = 0
	EMAIL_INVALID = 1
	EMAIL_DISPOSABLE = 2
	EMAIL_CATCHALL = 3
	EMAIL_UNKNOWN = 4
)

var NeverBounce *NeverBounceCli

type NeverBounceCli struct {
	ApiUrl      string
	ApiUsername string
	ApiPassword string
	AccessToken string

	TestMode bool
}

// Sets the API url on the client
func (n *NeverBounceCli) SetApiUrl(url string) {
	n.ApiUrl = url
}

// This function will get the client access token but also
// store it in the struct to be used in subsequent calls
func (n *NeverBounceCli) GetAccessToken() string {
	log.Println("Requesting access token")

	request, _ := http.NewRequest(
		"POST",
		n.ApiUrl+"/access_token",
		bytes.NewReader([]byte(url.Values{
			"grant_type": {"client_credentials"},
			"scope":      {"basic user"},
		}.Encode())),
	)

	request.SetBasicAuth(n.ApiUsername, n.ApiPassword)

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Panic(err)
	}

	var accessTokenResponse AccessTokenResponse
	decoder := json.NewDecoder(response.Body)

	if response.StatusCode != 200 {
		log.Panic("Client returned status: ", response.StatusCode)
		return ""
	}

	if err := decoder.Decode(&accessTokenResponse); err != nil {
		log.Panic(err)
	} else {
		n.AccessToken = accessTokenResponse.AccessToken
	}

	return accessTokenResponse.AccessToken
}

// Takes an email and verifies it
func (n *NeverBounceCli) VerifyEmail(email string) VerifyEmailResponse {
	log.Println("Verifying email ", email)

	response, err := http.PostForm(
		n.ApiUrl+"/single",
		url.Values{
			"access_token": {n.AccessToken},
			"email":        {email},
		})

	if err != nil {
		log.Panic(err)
	} else if response.StatusCode != 200 {
		log.Panic("Email verification returned status ", response.StatusCode)
	}

	var verifyEmailResponse VerifyEmailResponse

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&verifyEmailResponse); err != nil {
		log.Panic(err)
	}

	if !verifyEmailResponse.Success {
		if strings.Contains(verifyEmailResponse.Msg, "Authentication failed") {
			log.Println("Email verification failed: ", verifyEmailResponse.Msg)
			n.GetAccessToken()
			return n.VerifyEmail(email)
		} else {
			return verifyEmailResponse
		}
	}

	return verifyEmailResponse
}

func VerifyEmail(email string) VerifyEmailResponse {
	if NeverBounce == nil {
		log.Fatal("NeverBounce email not initiated.")
	}

	return NeverBounce.VerifyEmail(email)
}

func GetAccessToken() string {
	if NeverBounce == nil {
		log.Fatal("NeverBounce email not initiated.")
	}

	return NeverBounce.GetAccessToken()
}

func Init(neverbounce *NeverBounceCli) {
	NeverBounce = neverbounce

	// In case the API url has not been provided set the default one.
	// Used for testing
	if NeverBounce.TestMode {
		// Set url for default server
		NeverBounce.SetApiUrl(GetFakeService().URL)
	} else if NeverBounce.ApiUrl == "" {
		NeverBounce.SetApiUrl(DEFAULT_API_URL)
	}
}
