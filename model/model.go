package model

import (
	"context"
	"github.com/usagiga/Nurupoga/entity"
)

type ConfigModel interface {
	Load(configFilePath string) (config *entity.Config, err error)
}

type PingPongModel interface {
	WatchUpdatingMessage(ctx context.Context) (err error)
}
