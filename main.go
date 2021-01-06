package main

import (
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"io/ioutil"
	"log"
)

type twitterAuth struct {
	ConsumerKey       string `json: "consumerKey"`
	ConsumerSecret    string `json: "consumerSecret"`
	AccessToken       string `json: "accessToken"`
	AccessTokenSecret string `json: "accessTokenSecret"`
}

func main() {
	auth, err := ioutil.ReadFile("sample.json")
	if err != nil {
		log.Fatal(err)
	}

	anaconda.SetConsumerKey("consumerkey")
	anaconda.SetConsumerSecret("consumersecret")
	api := anaconda.NewTwitterApi("accesstoken", "accesstokensecret")

}
