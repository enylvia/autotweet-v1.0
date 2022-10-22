package main

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/rivo/uniseg"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var s = gocron.NewScheduler(time.UTC)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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
	// Welcome endpoint and ping to keep the app awake
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello your APP is running")
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
	s.Every(25).Minute().Do(Ping)
	// Start cron job
	s.StartAsync()
	fmt.Println("Cron job started, currently running jobs: ", len(s.Jobs()))
}

func Ping() {
	// Ping to keep the app awake
	http.Get("https://autotweets.herokuapp.com/")
}

func SendTweet(client *twitter.Client) {
	// Get Tweet data from function GetTweet
	tweet, err := GetTweet()
	if err != nil {
		log.Println(err)
	}
	// Format text that we want to send
	formatString := fmt.Sprintf(tweet.Data[0].AyaName + " - " + tweet.Data[0].Text)
	// Limit character
	count := uniseg.GraphemeClusterCount(formatString)
	if count > 275 {
		formatString = LimitCharacter(formatString, 275)
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
	min := 1
	max := 6236
	// Get random number
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(max-min) + min
	var data ResponseTweet
	resp, err := http.Get("https://api.myquran.com/v1/tafsir/quran/kemenag/id/" + strconv.Itoa(random))
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

func LimitCharacter(word string, limit int) string {
	// Limit Character
	reader := strings.NewReader(word)

	// Get the first 270 characters
	buff := make([]byte, limit)

	n, _ := io.ReadAtLeast(reader, buff, limit)

	if n != 0 {
		return string(buff) + "..."
	} else {
		return word
	}
}
