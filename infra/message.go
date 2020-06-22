package infra

import (
	"github.com/bwmarrin/discordgo"
)

type MessageInfraImpl struct {}

func NewMessageInfra() (messageInfra MessageInfra) {
	return &MessageInfraImpl{}
}

func (i *MessageInfraImpl) Reply(message string, ev *discordgo.MessageCreate, session *discordgo.Session) (err error) {
	destChan := ev.ChannelID
	_, err = session.ChannelMessageSend(destChan, message)
	if err != nil {
		return err
	}

	return nil
}
