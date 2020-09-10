package wire

import (
	"github.com/liampulles/banger/pkg/config"
	"github.com/liampulles/banger/pkg/library"
)

type Context interface {
	Config() config.Config
	LibraryService() library.Service
}

type ContextImpl struct {
	config         config.Config
	libraryService library.Service
}

var _ Context = &ContextImpl{}

func (c *ContextImpl) Config() config.Config {
	return c.config
}

func (c *ContextImpl) LibraryService() library.Service {
	return c.libraryService
}

func WireContext(configPath string) (Context, error) {
	configService := config.NewConfigServiceImpl(configPath)
	cfg, err := configService.Load()
	if err != nil {
		return nil, err
	}
	libraryService := library.NewService(cfg.LibraryRootPath())
	if err != nil {
		return nil, err
	}

	return &ContextImpl{
		config:         cfg,
		libraryService: libraryService,
	}, nil
}
