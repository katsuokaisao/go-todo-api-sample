//go:build wireinject
// +build wireinject

package serve

import (
	"github.com/google/wire"
	"github.com/katsuokaisao/gin-play/api"
	"github.com/katsuokaisao/gin-play/api/handler"
	"github.com/katsuokaisao/gin-play/domain"
	"github.com/katsuokaisao/gin-play/infra/rdb"
	"github.com/katsuokaisao/gin-play/usecase"
)

type Serve struct {
	Server *api.Server
}

func NewServe(server *api.Server) *Serve {
	return &Serve{
		Server: server,
	}
}

func SetUpServe(cfg *domain.Env) (*Serve, error) {
	wire.Build(
		NewServe,
		api.NewServer,
		domain.NewJWTParser,
		handler.NewTodoHandler,
		usecase.NewTodoUseCase,
		rdb.NewTodoRepository,
		rdb.NewRDB,
		provideDBEnv,
		provideAPIEnv,
	)
	return &Serve{}, nil
}

func provideAPIEnv(cfg *domain.Env) *domain.APIEnv {
	return cfg.API
}

func provideDBEnv(cfg *domain.Env) *domain.DBEnv {
	return cfg.DB
}
