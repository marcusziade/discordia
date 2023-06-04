# Discordia

Discordia is a simple yet powerful Discord bot, written in Go. It leverages the [discordgo](https://github.com/bwmarrin/discordgo) package to create a Discord bot that responds to certain keywords in messages. Currently, the bot can reply to "ping" with "Pong!" and can provide the current time when a user sends a message containing "time".

## Use Case

The bot can be used in any Discord server to automate specific responses. For instance, if a user wants to test the bot's responsiveness, they can type "ping", to which the bot will reply "Pong!". Similarly, if someone wants to know the current time, they just need to mention "time" in their message, and the bot will respond with the current time.

The potential applications of this bot are only limited by your imagination. You could extend its capabilities to include more commands, such as fetching weather data, responding with jokes, automating server management tasks, and much more.

## Getting Started

To start using Discordia, you need to have Go installed and a Discord bot token. Follow the steps below to get it running:

### Prerequisites

- Go (version 1.15 or higher)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/your_username_/discordia.git
```

2. Navigate to the repository:

```bash
cd discordia
```

3. Replace the dummy token with your actual Discord bot token in the main.go file:

```go
"Bot " + "YOUR_DISCORD_BOT_TOKEN",
```

4. Build and run the bot:

```bash
go build
./discordia
```

## Usage

Once the bot is running, invite it to your server and you can start interacting with it. The bot currently supports the following commands:

- `ping`: The bot will respond with "Pong!".
- `time`: The bot will respond with the current time.
