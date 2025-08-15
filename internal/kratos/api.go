package kratos

import (
	"context"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
)

type IKratos interface {
	Get(ctx context.Context, id uuid.UUID) (*ory.Identity, error)
}

type KratosService struct {
	client *ory.APIClient
}

func NewKratosService(url string) *KratosService {
	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			URL: url,
		},
	}
	return &KratosService{
		client: ory.NewAPIClient(configuration),
	}
}

func (s *KratosService) Get(ctx context.Context, id uuid.UUID) (*ory.Identity, error) {
	identity, _, err := s.client.IdentityAPI.GetIdentity(ctx, id.String()).Execute()
	if err != nil {
		return nil, err
	}
	return identity, nil
}
