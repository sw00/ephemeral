# ephemeral: automatically delete your old Tweets with AWS Lambda

**ephemeral** is a Twitter timeline grooming program that runs for pretty much free on AWS Lambda. The code is forked from Adam Drake's excellent [Harold](https://github.com/adamdrake/harold) butler-like bot and refactored for Lambda.

You can use ephemeral to automatically delete all tweets from your timeline that are older than a certain number of hours that you can choose. For instance, you can ensure that your tweets are deleted after one week (168h), or one day (24h).

The program will run once for each execution based on the trigger/schedule you set in AWS Lambda. It will delete up to 200 expired tweets (per-request limit set by Twitter's API) each run.

## Set up

You will need to [create a new Twitter application and generate API keys](https://apps.twitter.com/). The program assumes the following environment variables are set:

```sh
TWITTER_CONSUMER_KEY
TWITTER_CONSUMER_SECRET
TWITTER_ACCESS_TOKEN
TWITTER_ACCESS_TOKEN_SECRET
MAX_TWEET_AGE
```

`MAX_TWEET_AGE` expects a value of hours, such as: `MAX_TWEET_AGE = 72h`

Optionally, you can whitelist certain tweets and save them from deletion by setting the variable `WHITELIST` with the tweet's ID as the value. Find the ID as the string of numbers at the end of the tweet's URL, for example:

https://twitter.com/victoriadotdev/status/ `1052624100617785344`

You can also whitelist tweets using a substring contained in the tweet, for instance, the hashtag `#remember`.

Set one value to whitelist, or multiple values using the separator `:` like so:

```go
WHITELIST = 1052624100617785344:1052942396034609152:#remember
```

You can set these variables in AWS Lambda when you create your Lambda function. For a full walkthrough with screenshots for creating a Lambda function and uploading the code, read [this blog post](https://victoria.dev/blog/running-a-free-twitter-bot-on-aws-lambda/). Skip to setting environment variables at [this link](https://victoria.dev/blog/running-a-free-twitter-bot-on-aws-lambda/#2-configure-your-function).

## Upload with update.sh

This handy bash script is included to help you upload your function code to Lambda. It requires [AWS Command Line Interface](https://aws.amazon.com/cli/). To set up, do `pip install awscli` and follow these instructions for [Quick Configuration](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html).

## License

Copyright (C) 2018-2020 Victoria Drake

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
