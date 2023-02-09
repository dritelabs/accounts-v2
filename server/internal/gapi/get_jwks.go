package gapi

import (
	"context"

	pb "github.com/dritelabs/accounts/internal/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *AccountServer) GetJwks(ctx context.Context, in *pb.GetJwksRequest) (*anypb.Any, error) {
	return &anypb.Any{
		Value: []byte(s.jwks),
	}, nil
}
