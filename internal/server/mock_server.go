package server

import (
	"context"

	"github.com/Sacro/GolangTechTask/api"
	"github.com/Sacro/GolangTechTask/internal/models"
	"github.com/google/uuid"
)

var voteables []*models.Voteable

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

	voteables = append(voteables, &models.Voteable{
		UUID:     u.String(),
		Question: r.Question,
		Answers:  r.Answers,
	})

	return &api.CreateVoteableResponse{Uuid: u.String()}, nil
}

func (s MockServer) ListVoteables(context.Context, *api.ListVoteablesRequest) (*api.ListVoteablesResponse, error) {
	var v = make([]*api.Voteable, len(voteables))

	for i, voteable := range voteables {
		v[i] = &api.Voteable{
			Uuid:     voteable.UUID,
			Question: voteable.Question,
			Answers:  voteable.Answers,
		}
	}

	return &api.ListVoteablesResponse{
		Votables: v,
	}, nil
}

func (s MockServer) mustEmbedUnimplementedVotingServiceServer() {
	panic("not implemented")
}
