package server

import (
	"github.com/dritelabs/accounts/internal/config"
	pb "github.com/dritelabs/accounts/internal/proto/v1"
	"gorm.io/gorm"
)

type AccountServer struct {
	pb.UnimplementedAccountServer
	Config config.Config
	DB     *gorm.DB
}

type AccountServerConfig struct {
	Config config.Config
	DB     *gorm.DB
}

func New(c *AccountServerConfig) *AccountServer {
	return &AccountServer{
		Config: c.Config,
		DB:     c.DB,
	}
}