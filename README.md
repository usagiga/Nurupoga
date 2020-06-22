# Nurupoga

Once you call alias, the Discord bot put emoji asap.


## Installation

```sh
$ go get github.com/usagiga/Nurupoga
```


## Usage

### 1. Setup Discord App & Bot

Make new app [here](https://discordapp.com/developers/applications).
Then make new bot on the app.

Open it on your browser, then authorize the app in all servers which you wanted to sync emojis.

```
https://discordapp.com/oauth2/authorize?&client_id=APP_CLIENT_ID_HERE&scope=bot&permissions=10240
```

NOTICE : You shouldn't use PUBLIC BOT mode.


### 2. Configure

Write bot token and guild(a.k.a server)'s ID and aliases you very wanted into `config.json`.
Put it on a directory as you like.


### 3. Launch

Just run it on directory put `config.json` .
If posted alias messages, Nurupoga send substance.


## Dependencies

- Go (1.14 or higher)
- [github.com/bwmarrin/discordgo](https://github.com/bwmarrin/discordgo)


## License

MIT
