package v1

import (
	"github.com/SaidovZohid/market-project/config"
	"github.com/SaidovZohid/market-project/storage"
)

type handlerV1 struct {
	cfg *config.Config
	Storage storage.StorageI
}

type HandlerV1 struct {
	Cfg *config.Config
	Storage *storage.StorageI
}

func New(opt *HandlerV1) *handlerV1 {
	return &handlerV1{
		cfg: opt.Cfg,
		Storage: *opt.Storage,
	}
}

