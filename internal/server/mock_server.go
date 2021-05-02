package server

import (
	"context"

	"github.com/Sacro/GolangTechTask/api"
)

type MockServer struct {
	api.UnimplementedVotingServiceServer
}

func (s MockServer) CastVote(ctx context.Context, r *api.CastVoteRequest) (*api.CastVoteResponse, error) {
	panic("not implemented")
}

func (s MockServer) CreateVoteable(ctx context.Context, r *api.CreateVoteableRequest) (*api.CreateVoteableResponse, error) {
	panic("not implemented")
}

func (s MockServer) ListVoteables(context.Context, *api.ListVoteablesRequest) (*api.ListVoteablesResponse, error) {
	panic("not implemented")
}

func (s MockServer) mustEmbedUnimplementedVotingServiceServer() {
	panic("not implemented")
}
