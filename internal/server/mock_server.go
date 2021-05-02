package server

import (
	"context"

	"github.com/Sacro/GolangTechTask/api"
	"github.com/google/uuid"
)

var voteables []*api.Voteable

type MockServer struct {
	api.UnimplementedVotingServiceServer
}

func (s MockServer) CastVote(ctx context.Context, r *api.CastVoteRequest) (*api.CastVoteResponse, error) {
	panic("not implemented")
}

func (s MockServer) CreateVoteable(ctx context.Context, r *api.CreateVoteableRequest) (*api.CreateVoteableResponse, error) {
	u, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	voteables = append(voteables, &api.Voteable{
		Uuid:     u.String(),
		Question: r.Question,
		Answers:  r.Answers,
	})

	return &api.CreateVoteableResponse{Uuid: u.String()}, nil
}

func (s MockServer) ListVoteables(context.Context, *api.ListVoteablesRequest) (*api.ListVoteablesResponse, error) {

	return &api.ListVoteablesResponse{
		Votables: voteables,
	}, nil
}

func (s MockServer) mustEmbedUnimplementedVotingServiceServer() {
	panic("not implemented")
}
