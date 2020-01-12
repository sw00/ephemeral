package main

import (
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"log"

	"github.com/ChimeraCoder/anaconda"
	"github.com/aws/aws-lambda-go/lambda"

	"flag"
)

var (
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")
	maxTweetAge       = getenv("MAX_TWEET_AGE")
	whitelist         = getWhitelist()
)

// MyResponse for AWS SAM
type MyResponse struct {
	StatusCode string `json:"StatusCode"`
	Message    string `json:"Body"`
}

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}

func getWhitelist() []string {
	v := os.Getenv("WHITELIST")

	if v == "" {
		return make([]string, 0)
	}

	return strings.Split(v, ":")
}

func getTimeline(api *anaconda.TwitterApi) ([]anaconda.Tweet, error) {
	args := url.Values{}
	args.Add("count", "200")        // Twitter only returns most recent 20 tweets by default, so override
	args.Add("include_rts", "true") // When using count argument, RTs are excluded, so include them as recommended
	timeline, err := api.GetUserTimeline(args)
	if err != nil {
		return make([]anaconda.Tweet, 0), err
	}
	return timeline, nil
}

func isWhitelisted(id int64) bool {
	tweetID := strconv.FormatInt(id, 10)

	for _, w := range whitelist {
		if w == tweetID {
			return true
		}
	}
	return false
}

func deleteFromTimeline(api *anaconda.TwitterApi, ageLimit time.Duration) {
	timeline, err := getTimeline(api)
	if err != nil {
		log.Print("could not get timeline", err)
	}

	for _, t := range timeline {
		createdTime, err := t.CreatedAtTime()
		if err != nil {
			log.Print("could not parse time ", err)
		} else {
			if time.Since(createdTime) > ageLimit && !isWhitelisted(t.Id) {
				_, err := api.DeleteTweet(t.Id, true)
				log.Print("DELETED ID ", t.Id)
				log.Print("TWEET ", createdTime, " - ", t.Text)
				if err != nil {
					log.Print("failed to delete: ", err)
				}
			}
		}
	}
	log.Print("no more tweets to delete")

}

func ephemeral() (MyResponse, error) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	api.SetLogger(anaconda.BasicLogger)

	h, _ := time.ParseDuration(maxTweetAge)

	deleteFromTimeline(api, h)

	return MyResponse{
		Message:    "no more tweets to delete",
		StatusCode: "200",
	}, nil
}

func main() {
	plain := flag.Bool("plain", false, "Specify to run in plain script mode instead of inside of lambda. In this mode, the application will behave like a simple CLI app instead of being bound to AWS.")
	flag.Parse()

	if *plain {
		_, err := ephemeral()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		lambda.Start(ephemeral)
	}
}
