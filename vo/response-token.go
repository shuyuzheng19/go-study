package vo

type ResponseToken struct {
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
