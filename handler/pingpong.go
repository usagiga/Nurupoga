package handler

import (
	"context"
	"github.com/usagiga/Nurupoga/model"
	"log"
	"os"
	"os/signal"
)

type PingPongCommandImpl struct {
	pingPongModel model.PingPongModel
}

func NewPingPongCommand(pingPongModel model.PingPongModel) (pingPongCommand PingPongCommand) {
	return &PingPongCommandImpl{
		pingPongModel: pingPongModel,
	}
}

func (c *PingPongCommandImpl) RunSyncService() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	errCh := make(chan error)

	// Run
	go func(){
		err := c.pingPongModel.WatchUpdatingMessage(ctx)
		if err != nil {
			errCh <- err
		}
	}()

	// Wait ^C or to raise errors
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	log.Println("Nurupoga(Stream-Mode) is now running... Press ^C to quit")

	select {
	case <-interrupt:
		log.Println("Keyboard Interrupt is detected. Quiting...")
		cancel()
	case err := <-errCh:
		log.Fatal("Error raised: ", err)
	}
}


