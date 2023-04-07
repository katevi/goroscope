# goroscope ![build](https://github.com/katevi/goroscope/actions/workflows/go.yml/badge.svg?style=shield??branch=master)[![Docker Image Size](https://badgen.net/docker/size/katevi/goroscope/?icon=docker&label=image%20size)](https://hub.docker.com/r/katevi/goroscope)

This application is a simple chat bot for Telegram that will return funny horoscopes. You can access it by the following [link](https://t.me/goroscope_katevi_bot).

## Tools
- bot is made with Go 1.20
- bot was integrated into Telegram sybsystem via [botFather](https://telegram.me/BotFather)
- bot functions are developed using [Telegram Bot API](https://core.telegram.org/bots/api)
- bot deployed to the Cloud via [Yandex Compute Cloud](https://cloud.yandex.ru/services/compute)

## Quick How-to

### Deploy bot locally
1. Clone this repository
2. Install Go 1.20
3. Register blank bot via [botFather](https://telegram.me/BotFather) and get token to access the HTTP API
4. Create the `.env` file in the root of repository with the following content:
```
TELEGRAM_BOT_TOKEN=<your token>
```
5. Run command `go run main.go` in the root of repository

### Deploy bot via creating Docker image
1. Clone this repository
2. Register blank bot via [botFather](https://telegram.me/BotFather) and get token to access the HTTP API
3. Create the `.env` file in the root of repository with the following content:
```
TELEGRAM_BOT_TOKEN=<your token>
```
4. Run following commands in the root of repository
```
docker build -t your-favorite-tag
docker run -v <abolute-path-to-env-file>/.env:/app/.env your-favorite-tag
```

### Deploy bot using public DockerHub image
1. Clone this repository
2. Register blank bot via [botFather](https://telegram.me/BotFather) and get token to access the HTTP API
3. Create the `.env` file in the root of repository with the following content:
```
TELEGRAM_BOT_TOKEN=<your token>
```
4. Run following command in the root of repository
```
docker run -v <abolute-path-to-env-file>/.env:/app/.env katevi/goroscope
```

You can look to this public docker image in more details at the DockerHub repository using the following [link](https://hub.docker.com/r/katevi/goroscope) which is also available via README badge.