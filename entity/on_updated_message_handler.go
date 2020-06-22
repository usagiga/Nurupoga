package entity

import "github.com/bwmarrin/discordgo"

// OnUpdatedMessageHandler is a handler to do something on updated messages.
type OnUpdatedMessageHandler func(message string, ev *discordgo.MessageCreate, session *discordgo.Session) (err error)

