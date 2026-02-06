package client

import (
	"context"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/internal"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/internal/logger"
	"gitlab.com/shaninalex/lumna/app/internal/persistence"
	"gorm.io/gorm"
)

type Client struct {
	ctx    context.Context
	logger logger.Logger
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
	return persistence.GetDB(s.ctx)
}

func (s *Client) Logger() logger.Logger {
	return s.logger
}

func NewClientForCLI(cmd *cobra.Command) (*Client, error) {
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		return nil, err
	}

	cnf := config.ReadConfig(configPath)
	// NOTE: context is not for data, it's about lifespan
	ctx := context.WithValue(cmd.Context(), internal.ContextConfig, cnf)
	ctx = context.WithValue(ctx, internal.ContextDB, persistence.Connect(cnf.Database.Url))

	client := &Client{
		ctx:    ctx,
		logger: logger.ProvideLogger(ctx),
	}
	return client, nil
}
