package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Token struct {
	ConsumerKey    string
	ConsumerSecret string
	ApiKey         string
	ApiSecret      string
}

func main() {

	fmt.Println("Auto Tweet V1.0")
	fmt.Println("=======================================")
	token := Token{
		ConsumerKey:    "YOUR_CONSUMER_KEY",
		ConsumerSecret: "YOUR_CONSUMER_SECRET",
		ApiKey:         "YOUR_API_KEY",
		ApiSecret:      "YOUR_API_SECRET",
	}
	client, err := getClient(&token)

	if err != nil {
		log.Println("Account not found")
		log.Println(err)
	}
	//CHANGE THIS IF YOU HAVE YOUR OWN DB OR TXT FILE
	file, err := os.Open("db/tweet.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tweet, resp, err := client.Statuses.Update(scanner.Text(), nil)
		if err != nil {
			log.Println(err)
		}
		var status = resp.StatusCode
		var text = tweet.Text
		log.Println("Status Code : ", status)
		log.Println("Tweet : ", text)

		//CRON JOB CHANGE A VALUE TO SET UR TIMER
		time.Sleep(time.Minute * 1)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func getClient(tkn *Token) (*twitter.Client, error) {
	config := oauth1.NewConfig(tkn.ConsumerKey, tkn.ConsumerSecret)
	tokenapi := oauth1.NewToken(tkn.ApiKey, tkn.ApiSecret)

	httpClient := config.Client(oauth1.NoContext, tokenapi)
	client := twitter.NewClient(httpClient)

	verifcek := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, err := client.Accounts.VerifyCredentials(verifcek)

	if err != nil {
		return nil, err
	}
	var regex, _ = regexp.Compile(`[a-z]+`)
	var info = regex.FindAllString(user.Name, -1)
	data := strings.Join(info, " ")
	fmt.Println("Account name : ", data)

	return client, nil
}
