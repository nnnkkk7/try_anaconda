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

	var twitterauth twitterAuth
	if err := json.Unmarshal(auth, &twitterauth); err != nil {
		log.Fatal(err)
	}

	anaconda.SetConsumerKey("consumerkey")
	anaconda.SetConsumerSecret("consumersecret")
	api := anaconda.NewTwitterApi("accesstoken", "accesstokensecret")

	searchResult, _ := api.GetSearch("Go言語", nil)
	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}

}
