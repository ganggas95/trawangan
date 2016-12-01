package job

import (
	"flag"
	"github.com/antonholmquist/jason"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
)

type TwitHandler struct {
}

func (t TwitHandler) GetUrlTwit() (string, *oauth.RequestToken, *oauth.Consumer) {
	var consumerKey *string = flag.String(
		"consumerkey",
		"6etLPfdmROyD3VZI6QC7kUNLw",
		"Consumer Key from Twitter. See: https://dev.twitter.com/apps/new")

	var consumerSecret *string = flag.String(
		"consumersecret",
		"vemENVtuEnEzVBbsfLdb1JU5LhNDujLKXuDTlL3YCUb5s8LDXl",
		"Consumer Secret from Twitter. See: https://dev.twitter.com/apps/new")

	flag.Parse()

	c := oauth.NewConsumer(
		*consumerKey,
		*consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})

	c.Debug(true)
	requestToken, u, err := c.GetRequestTokenAndUrl("https://trawangan.herokuapp.com/verifytwit")
	if err != nil {
		log.Fatal(err)
	}
	return u, requestToken, c
}
func (t TwitHandler) VerifyTwit(code string) *jason.Object {
	_, requestToken, c := t.GetUrlTwit()
	accessToken, err := c.AuthorizeToken(requestToken, code)
	if err != nil {
		log.Fatal(err)
	}

	client, err := c.MakeHttpClient(accessToken)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Get(
		"https://api.twitter.com/1.1/account/verify_credentials.json?include_email=true")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	user, _ := jason.NewObjectFromBytes(bits)
	return user
}
