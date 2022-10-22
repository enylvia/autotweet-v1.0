package main

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/go-co-op/gocron"
	"github.com/rivo/uniseg"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var s = gocron.NewScheduler(time.UTC)

func main() {
	// Load .env file
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	fmt.Println("Auto Tweet From API V1.0")
	fmt.Println("=======================================")
	// Get Token From .env
	token := Token{
		ConsumerKey:    os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
		ApiKey:         os.Getenv("API_KEY"),
		ApiSecret:      os.Getenv("API_SECRET"),
	}
	// Login to Twitter using the credentials
	client, err := getClient(&token)
	if err != nil {
		log.Println(err)
	}
	// Welcome endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello your App is running")
	})
	// Endpoint for start cron job
	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		RunCronJob(client)
	})
	// Endpoint for stop cron job
	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		StopCronJob()
	})
	// Run server from port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)

}
func StopCronJob() {
	// Stop cron job
	// Clear all scheduler jobs
	s.Clear()
	fmt.Println("Jobs cleared, currently running jobs: ", len(s.Jobs()))
}
func RunCronJob(client *twitter.Client) {
	// Start cron job to send tweet every 1 hour
	s.Every(1).Hour().Do(SendTweet, client)
	// Start cron job
	s.StartAsync()
	fmt.Println("Cron job started, currently running jobs: ", len(s.Jobs()))
}

func SendTweet(client *twitter.Client) {
	// Get Tweet data from function GetTweet
	tweet, err := GetTweet()
	if err != nil {
		log.Println(err)
	}
	// Format text that we want to send
	formatString := fmt.Sprintf(tweet.Acak.Id.Teks+"(QS:%s: %s)", tweet.Acak.Id.Surat, tweet.Acak.Id.Ayat)
	// Check if the text is more than 270 characters
	count := uniseg.GraphemeClusterCount(formatString)
	if count > 270 {
		fmt.Println("Tweet is too long, skipping...")
	}
	// Send tweet to Twitter
	send, resp, err := client.Statuses.Update(formatString, nil)
	if err != nil {
		log.Println(err)
	}
	// Check if the tweet is sent successfully
	fmt.Printf("TWEETED: %+v , %v\n", send.Text, resp.Status)
}
func GetTweet() (ResponseTweet, error) {
	// Get Data From API
	var data ResponseTweet
	resp, err := http.Get("https://api.banghasan.com/quran/format/json/acak")
	if err != nil {
		return ResponseTweet{}, err
	}
	defer resp.Body.Close()
	// Decode data from API
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return ResponseTweet{}, err
	}
	return data, nil
}

// Get Client Twitter
func getClient(tkn *Token) (*twitter.Client, error) {
	// Get Token From .env file
	config := oauth1.NewConfig(tkn.ConsumerKey, tkn.ConsumerSecret)
	tokenapi := oauth1.NewToken(tkn.ApiKey, tkn.ApiSecret)
	// Get Client Twitter
	httpClient := config.Client(oauth1.NoContext, tokenapi)
	client := twitter.NewClient(httpClient)

	// Check if the client is valid
	verifcek := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// Get user information from Twitter
	user, _, err := client.Accounts.VerifyCredentials(verifcek)

	if err != nil {
		return nil, err
	}
	// Print user information
	var regex, _ = regexp.Compile(`[a-z]+`)

	// Get username from Twitter
	var info = regex.FindAllString(user.ScreenName, -1)
	data := strings.Join(info, " ")
	fmt.Println("Username : @" + data)
	fmt.Println("Current Tweet : ", user.StatusesCount)
	fmt.Println("Follower : ", user.FollowersCount)
	fmt.Println("=======================================")
	return client, nil
}
