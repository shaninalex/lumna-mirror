package client

import (
	"context"
	"database/sql"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/global"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/services"
)

type Client struct {
	ctx context.Context

	UserManager services.UserManager
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

func (s *Client) DBConnect() *sql.DB {
	cnf := s.Config()
	conn, err := sql.Open("sqlite3", cnf.Database.Url)
	if err != nil {
		panic(err)
	}

	s.ctx = context.WithValue(s.ctx, global.ContextDB, conn)
	return conn
}

func NewClient(cmd *cobra.Command) (*Client, error) {
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		return nil, err
	}
	cnf := config.ReadConfig(configPath)

	return &Client{
		ctx:         context.WithValue(cmd.Context(), global.ContextConfig, cnf),
		UserManager: services.NewUserService(),
	}, nil
}
