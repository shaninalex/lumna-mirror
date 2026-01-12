package client

import (
	"context"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/internal"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gorm.io/gorm"
)

type Client struct {
	ctx context.Context
}

func (s *Client) Context() context.Context {
	return s.ctx
}

func (s *Client) Config() *config.Config {
	cnf, ok := s.ctx.Value(internal.ContextConfig).(*config.Config)
	if !ok {
		panic("config not found in context")
	}
	return cnf
}

func (s *Client) DB() *gorm.DB {
	conn, ok := s.ctx.Value(internal.ContextDB).(*gorm.DB)
	if !ok {
		panic("db not found in context")
	}
	return conn
}

func (s *Client) initDB() {
	cnf := s.Config()
	s.ctx = context.WithValue(s.ctx, internal.ContextDB, db.Connect(cnf.Database.Url))
}

func NewClient(cmd *cobra.Command) (*Client, error) {
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		return nil, err
	}

	cnf := config.ReadConfig(configPath)
	ctx := context.WithValue(cmd.Context(), internal.ContextConfig, cnf)
	ctx = context.WithValue(ctx, internal.ContextDB, db.Connect(cnf.Database.Url))

	client := &Client{
		ctx: ctx,
	}
	return client, nil
}
