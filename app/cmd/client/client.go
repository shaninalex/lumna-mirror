package client

import (
	"context"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/pkg/global"
	"gorm.io/gorm"
)

type Client struct {
	ctx context.Context
}

func (s *Client) Context() context.Context {
	return s.ctx
}

func (s *Client) Config() *config.Config {
	cnf, ok := s.ctx.Value(global.ContextConfig).(*config.Config)
	if !ok {
		panic("config not found in context")
	}
	return cnf
}

func (s *Client) DB() *gorm.DB {
	conn, ok := s.ctx.Value(global.ContextDB).(*gorm.DB)
	if !ok {
		panic("db not found in context")
	}
	return conn
}

func (s *Client) initDB() {
	cnf := s.Config()
	s.ctx = context.WithValue(s.ctx, global.ContextDB, db.Connect(cnf.Database.Url))
}

func NewClient(cmd *cobra.Command) (*Client, error) {
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		return nil, err
	}

	cnf := config.ReadConfig(configPath)

	ctx := context.WithValue(cmd.Context(), global.ContextConfig, cnf)
	ctx = context.WithValue(ctx, global.ContextDB, db.Connect(cnf.Database.Url))

	client := &Client{
		ctx: ctx,
	}
	return client, nil
}
