# DESCRIPTION
[]: # Title: README.md

This Auto Tweet allows you to do automatically tweet. I made this because I wanted to tweet
automatically for my Twitter Bot.

## Prerequisites

Before you begin, ensure you have met the following requirements:
* You have installed the latest version of Golang
* You have a Windows/Linux/Mac machine. State which OS is supported/which is not.
* Some knowledge of Golang
* Some Text Editor

## Installing <Auto Tweet>

```
git clone https://github.com/enylvia/autotweet-v1.git
```
## Using <Auto Tweet>

To use <Auto Tweet>, follow these steps:
Do this first : 
```
Create File and name it ".env"
    And put this code :
    CONSUMER_KEY=YOUR_CONSUMER_KEY
    CONSUMER_SECRET=YOUR_CONSUMER_SECRET
    API_KEY=YOUR_API_KEY
    API_SECRET=YOUR_API_SECRET
```

You can get this from  : https://apps.twitter.com/

And you can change the tweet from function get tweet.
## HOW TO RUN

After you edit ConsumerKey,etc.. 
```
go mod tidy 
go build . 
OR 
go run main.go
```

## Contact

If you want to contact me you can reach me at <addityap@hotmail.com>.
Twitter : https://twitter.com/additya_pp
