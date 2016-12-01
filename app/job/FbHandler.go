package job

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"net/http"
	"net/url"
)

type FbHandler struct {
}

var FACEBOOK = &oauth2.Config{
	ClientID:     "1232747110132777",
	ClientSecret: "4daa48f88b21adf49913e3d9abed3faa",
	Scopes:       []string{"email", "user_birthday", "user_location", "user_about_me"},
	Endpoint:     facebook.Endpoint,
	RedirectURL:  "https://trawangan.herokuapp.com/authfb",
}
var FACEBOOK2 = &oauth2.Config{
	ClientID:     "1232747110132777",
	ClientSecret: "4daa48f88b21adf49913e3d9abed3faa",
	Scopes:       []string{"email", "user_birthday", "user_location", "user_about_me"},
	Endpoint:     facebook.Endpoint,
	RedirectURL:  "https://trawangan.herokuapp.com/loginwithfb",
}

func (fb FbHandler) GetUrlFb() string {
	authUrl := FACEBOOK.AuthCodeURL("")
	return authUrl
}
func (fb FbHandler) GetUrlLoginFb() string {
	authUrl := FACEBOOK2.AuthCodeURL("")
	return authUrl
}

func (fb FbHandler) GetTokenFb(code string) *oauth2.Token {
	var tok *oauth2.Token
	tok, err := FACEBOOK.Exchange(nil, code)
	if err != nil {
		panic(err)
	}
	return tok
}

func (fb FbHandler) GetResponseFb(tok *oauth2.Token) *http.Response {
	response, err := http.Get("https://graph.facebook.com/me?access_token=" + url.QueryEscape(tok.AccessToken))
	if err != nil {
		panic(err)
	}
	return response
}
