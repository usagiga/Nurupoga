package infra

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/usagiga/Nurupoga/entity"
)

type MessageStreamInfraImpl struct {
	cred entity.Credential
}

func NewMessageStreamInfra(cred entity.Credential) (messageStreamInfra MessageStreamInfra) {
	return &MessageStreamInfraImpl{cred: cred}
}

func (i *MessageStreamInfraImpl) WatchUpdatingMessage(ctx context.Context, handler entity.OnUpdatedMessageHandler) (err error) {
	errCh := make(chan error)
	token := i.cred.GetBearerToken()
	onUpdateMessage := func(s *discordgo.Session, m *discordgo.MessageCreate) {
		message := m.Content

		// Ignore all messages created by the bot itself
		// This isn't required in this specific example but it's a good practice.
		if m.Author.ID == s.State.User.ID {
			return
		}

		// Fire OnUpdatedMessageHandler
		err := handler(message, m, s)
		if err != nil {
			errCh <- err
		}
	}

	// Listen the registered event.
	discord, err := discordgo.New(token)
	if err != nil {
		errCh <- err
		return
	}

	discord.AddHandler(onUpdateMessage)
	err = discord.Open()
	if err != nil {
		errCh <- err
		return
	}
	defer discord.Close()

	// Wait for interrupting or raising errors.
	select {
	case err = <-errCh:
		return err
	case <-ctx.Done():
		return nil
	}
}
