# goroscope ![build](https://github.com/katevi/goroscope/actions/workflows/go.yml/badge.svg?style=shield??branch=master)[![Docker Image Size](https://badgen.net/docker/size/katevi/goroscope/?icon=docker&label=image%20size)](https://hub.docker.com/r/katevi/goroscope)[![Docker Pulls](https://badgen.net/docker/pulls/katevi/goroscope?icon=docker&label=pulls)](https://hub.docker.com/r/katevi/goroscope/)

This application is a simple chat bot for Telegram that will return funny horoscopes. You can access it by the following [link](https://t.me/goroscope_katevi_bot).

## Tools
- bot is made with Go 1.20
- bot uses [Redis](https://redis.io/) for storing data
- bot integrated into Telegram sybsystem via [botFather](https://telegram.me/BotFather)
- bot functions are developed using [Telegram Bot API](https://core.telegram.org/bots/api)
- bot deployed to the Cloud via [Yandex Compute Cloud](https://cloud.yandex.ru/services/compute)

## Quick How-to

1. Clone this repository
2. Register blank bot via [botFather](https://telegram.me/BotFather) and get token to access the HTTP API
3. Run following command in the root of repository
```
TELEGRAM_BOT_TOKEN=<your-token> docker compose up
```
You can look to this public docker image at the DockerHub repository using the following [link](https://hub.docker.com/r/katevi/goroscope) which is also available via README badge.

## Development guide
1. Clone this repository
2. Install Go 1.20
3. Register blank bot via [botFather](https://telegram.me/BotFather) and get token to access the HTTP API
4. Update `docker-compose.yaml` to build goroscope from local directory instead of using public DockerHub image
5. Develop amazing features to goroscope :)
6. Run following command in the root of repository to start goroscope locally:
```
TELEGRAM_BOT_TOKEN=<your-token> docker compose up
```

Docker compose starts container with goroscope app and Redis container for storing goroscope's data.