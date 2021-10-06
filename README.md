# Auto Tweet V1.0

Project name is a Auto Tweet V1.0 that allows you to do automatically tweet.

## Prerequisites

Before you begin, ensure you have met the following requirements:
* You have installed the latest version of Golang

## Installing <Auto Tweet V1.0>

```
git clone https://github.com/enylvia/autotweet-v1.git
```
## Using <Auto Tweet V1.0>

To use <Auto Tweet V1.0>, follow these steps:
Change this first : 
```
    		ConsumerKey:    "YOUR_CONSUMER_KEY",
		ConsumerSecret: "YOUR_CONSUMER_SECRET",
		ApiKey:         "YOUR_API_KEY",
		ApiSecret:      "YOUR_API_SECRET",
```

You can get this from  : https://apps.twitter.com/

And you can change the txt file on db folder.
Then you need to edit this : 
```
file, err := os.Open("db/tweet.txt")
```
## HOW TO RUN

After you edit ConsumerKey,etc.. 
```
go run main.go
```

## Contact

If you want to contact me you can reach me at <adityaperm06@gmail.com>.
