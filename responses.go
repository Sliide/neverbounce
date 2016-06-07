package neverbounce

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Expires     int    `json:"expires"`
}

type VerifyEmailResponse struct {
	Success       bool   `json:"success"`
	Result        int    `json:"result"`
	ResultDetails int    `json:"result_details"`
	Msg           string `json:"msg"`
}
