# TelegramNewsbot
A bot for telegram which sends news for a given country into a channel

# Demo
Germany news: [@germany_news_channel](https://t.me/germany_news_channel)

# Get started
Compile it yourself or use the [docker image](https://hub.docker.com/r/jojii/newsbot)<br>
Configure the bot<br>
Add the bot to your channel and give him rights to send messages

### Compile
```bash
go mod download &&
go build -o newsbot
```

### Docker setup
```bash
export TAG=$(curl -s -S "https://registry.hub.docker.com/v2/repositories/jojii/newsbot/tags/" | jq '.results[]["name"]' -r | sed -n 1p)
docker pull jojii/newsbot:$TAG
docker run -d --name newsBot --restart=unless-stopped -v `pwd`/data:/app/data -v /etc/localtime:/etc/localtime:ro jojii/newsbot:$TAG
```

## Config
`TelegramToken` the telegram bot token. If you don't have one, create a new bot with @BotFather<br>
`NewsAPIToken` The API token for the [newsAPI](https://newsapi.org/)<br>
`ChannelID` The ID of the channel<br>
`StoreFile` A file containing some data needed for the bot to run correctly<br>
`Country` Selector for the news (by country)<br>