package entity

type Credential struct {
	BotToken string `json:"bot_token"`
	GuildID string `json:"guild_id"`
}

// GetBearerToken generates the body of `Authorization` HTTP header.
func (cred *Credential) GetBearerToken() string {
	return "Bot " + cred.BotToken
}
