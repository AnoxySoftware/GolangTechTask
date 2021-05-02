package server

import (
	"context"

	"github.com/Sacro/GolangTechTask/api"
)

type DynamodbServer struct {
	api.UnimplementedVotingServiceServer
}

func (s DynamodbServer) CastVote(ctx context.Context, r *api.CastVoteRequest) (*api.CastVoteResponse, error) {
	panic("not implemented")
}

func (s DynamodbServer) CreateVoteable(ctx context.Context, r *api.CreateVoteableRequest) (*api.CreateVoteableResponse, error) {
	panic("not implemented")
}

func (s DynamodbServer) ListVoteables(context.Context, *api.ListVoteablesRequest) (*api.ListVoteablesResponse, error) {
	panic("not implemented")
}
