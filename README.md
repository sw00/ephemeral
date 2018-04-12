# ephemeral: automatically delete your old Tweets with AWS Lambda

**ephemeral** is a Twitter timeline grooming program that runs for pretty much free on AWS Lambda. The code is forked from Adam Drake's excellent [Harold](https://github.com/adamdrake/harold) butler-like bot.

You can use ephemeral to automatically delete all tweets from your timeline that are older than a certain number of hours that you can choose. For instance, you can ensure that your tweets are deleted after one week (168h), or one day (24h).

# Twitter API

You will need to [create a new Twitter application and generate API keys](https://apps.twitter.com/). The program assumes the following environment variables are set:

```
TWITTER_CONSUMER_KEY
TWITTER_CONSUMER_SECRET
TWITTER_ACCESS_TOKEN
TWITTER_ACCESS_TOKEN_SECRET
MAX_TWEET_AGE
```

`MAX_TWEET_AGE` expects a value of hours, such as: `MAX_TWEET_AGE=72h`

You can set these variables in AWS Lambda when you create your Lambda function. For a full walkthrough with screenshots on creating a Lambda function and uploading the code, read [this blog post](https://vickylai.com/verbose/aws-lambda/). Skip to setting environment variables at [this link](https://vickylai.com/verbose/aws-lambda/#2-configure-your-function).

The program will run once for each execution based on the trigger/schedule you set in AWS Lambda.