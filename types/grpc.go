package types

import (
	"context"

	"google.golang.org/grpc"
	plugin "github.com/hashicorp/go-plugin"
	sdk "github.com/cosmos/cosmos-sdk/types"
	codec "github.com/cosmos/cosmos-sdk/codec/types"
)

type GRPCClient struct {
	broker *plugin.GRPCBroker
	client ModuleClient
}

func (m *GRPCClient) Handler(stateServer StateServer) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		var s *grpc.Server
		serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
			s = grpc.NewServer(opts...)
			RegisterStateServer(s, stateServer)
			return s
		}

		brokerID := m.broker.NextId()
		go m.broker.AcceptAndServe(brokerID, serverFunc)

		resp, err := m.client.Handle(context.Background(), &RequestHandle {
			Msg: &codec.Any{msg},
			Context: ContextToProto(ctx),
		})

		s.Stop()

		if err != nil {
			return nil, err
		}

		return resp.Result, nil // TODO: use proper err
	}
}
