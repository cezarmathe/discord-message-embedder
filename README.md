# discord-message-embedder

A discord bot that sends embedded messages for you by parsing a JSON string.

## Setup

There are two(one, executables are coming soon) ways to install the bot:

1. Install and/or run the bot using the `go` tool(this requires the `go` tool and a proper `GOPATH`)

- Using `go get` and running the executable directly(this requires your `GOBIN` to be added in your `PATH`)

```bash
go get github.com/cezarmathe/discord-message-embedder
DISC_MSG_EMBEDDER_TOKEN=my-token discord-message-embedder
```

- Using `git clone`

```bash
git clone git://github.com/cezarmathe/discord-message-embedder.git
cd discord-message-embedder
# use this if you're lazy
DISC_MSG_EMBEDDER_TOKEN=my-token go run main.go
# use this if you don't have GOBIN in your PATH
go build
chmod +x discord-message-embedder
DISC_MSG_EMBEDDER_TOKEN=my-token ./discord-message-embedder
# use this if you have GOBIN in your path(though it would've been easier to try the first installation method)
go install
DISC_MSG_EMBEDDER_TOKEN=my-token discord-message-embedder
```

2. Downloading the executable

**COMING SOON**

## How to use

1. Invite the bot to your server

The proper invite link for this bot is `https://discordapp.com/oauth2/authorize?client_id=your_app_id&scope=bot&permissions=55296`. You will have to self-host the bot yourself, or run it every time you need to send an embedded message.

2. Bot information

This bot uses the prefix `!`, if you'd like to change the prefix, change the variable `commandPrefix = "!"` from _main.go_ to something else.

3. Embedding messages

To get a sample JSON that is required as input, send:

> !embed

You can also remove any fields from JSON, as needed.

To send an embedded message from a JSON, send:

> !embed channel_id my_json_string