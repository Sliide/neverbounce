package neverbounce

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

func GetFakeService() *httptest.Server {

	log.Println("Starting fake service.")

	mux := http.NewServeMux()
	mux.HandleFunc("/access_token", func(w http.ResponseWriter, r *http.Request) {

		at := AccessTokenResponse{
			AccessToken: "thisisatoken",
			Expires:     360,
		}

		json, _ := json.Marshal(at)
		w.Write(json)
	})

	mux.HandleFunc("/single", func(w http.ResponseWriter, r *http.Request) {
		// Function to simulate single email check

		body, _ := ioutil.ReadAll(r.Body)
		values, err := url.ParseQuery(string(body))

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		accessToken := values.Get("access_token")
		email := values.Get("email")

		var ve VerifyEmailResponse

		if accessToken != "thisisatoken" {
			ve = VerifyEmailResponse{
				Success: false,
				Msg:     "Authentication failed",
			}
		}

		if email == "" {
			ve = VerifyEmailResponse{
				Success: false,
				Msg:     "Missing required parameter 'email'",
			}
		}

		if strings.Contains(email, "@valid.com") {
			ve = VerifyEmailResponse{
				Success: true,
				Result:  EMAIL_VALID,
			}

		} else if strings.Contains(email, "@invalid.com") {
			ve = VerifyEmailResponse{
				Success: true,
				Result:  EMAIL_INVALID,
			}

		} else if strings.Contains(email, "@disposable.com") {
			ve = VerifyEmailResponse{
				Success: true,
				Result:  EMAIL_DISPOSABLE,
			}

		} else if strings.Contains(email, "@catchall.com") {
			ve = VerifyEmailResponse{
				Success: true,
				Result:  EMAIL_CATCHALL,
			}

		} else if strings.Contains(email, "@unknown.com") {
			ve = VerifyEmailResponse{
				Success: true,
				Result:  EMAIL_UNKNOWN,
			}
		} else {
			ve = VerifyEmailResponse{
				Success: true,
				Result:  EMAIL_INVALID,
			}
		}

		json, _ := json.Marshal(ve)

		w.Write(json)
	})

	return httptest.NewServer(mux)
}
