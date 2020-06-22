package infra

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/usagiga/Nurupoga/entity"
)

// MessageInfra is to send messages specified server.
type MessageInfra interface {
	Reply(message string, ev *discordgo.MessageCreate, session *discordgo.Session) (err error)
}

// MessageStreamInfra treats the events on Discord API server.
// Particularly about updating messages.
type MessageStreamInfra interface {
	// WatchUpdatingMessage waits updating message on specified servers
	// and notifies about it to the handler.
	WatchUpdatingMessage(ctx context.Context, handler entity.OnUpdatedMessageHandler) (err error)
}
