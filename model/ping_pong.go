package model

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/usagiga/Nurupoga/entity"
	"github.com/usagiga/Nurupoga/infra"
	"strings"
)

type PingPongModelImpl struct {
	config             *entity.Config
	messageInfra       infra.MessageInfra
	messageStreamInfra infra.MessageStreamInfra
}

func NewPingPongModel(
	config *entity.Config,
	messageInfra infra.MessageInfra,
	messageStreamInfra infra.MessageStreamInfra,
) (model PingPongModel) {
	return &PingPongModelImpl{
		config:             config,
		messageInfra:       messageInfra,
		messageStreamInfra: messageStreamInfra,
	}
}

func (m *PingPongModelImpl) WatchUpdatingMessage(ctx context.Context) (err error) {
	var onUpdateMessage entity.OnUpdatedMessageHandler
	onUpdateMessage = func(message string, ev *discordgo.MessageCreate, session *discordgo.Session) (err error) {
		for _, alias := range m.config.Aliases {
			if strings.Index(message, alias.Alias) != -1 {
				// Reply
				err = m.messageInfra.Reply(alias.Substance, ev, session)
				if err != nil {
					return err
				}

				break
			}
		}

		return nil
	}

	err = m.messageStreamInfra.WatchUpdatingMessage(ctx, onUpdateMessage)
	if err != nil {
		return err
	}

	return nil
}
