package server

import (
	"context"

	"github.com/Sacro/GolangTechTask/api"
	"github.com/Sacro/GolangTechTask/internal/models"
	"github.com/apex/log"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/guregu/dynamo"
)

const tableName = "Voteables"

type DynamodbServer struct {
	api.UnimplementedVotingServiceServer

	table dynamo.Table
}

func NewDynamodbServer(db *dynamo.DB) *DynamodbServer {
	err := db.CreateTable(tableName, models.Voteable{}).Run()
	if err != nil {
		// If ResourceInUseException is thrown our table already exists
		if err.(awserr.Error).Code() != "ResourceInUseException" {

			log.WithError(err).Info("creating table")
		}
	}

	table := db.Table(tableName)

	return &DynamodbServer{
		table: table,
	}
}

func (s DynamodbServer) CastVote(ctx context.Context, r *api.CastVoteRequest) (*api.CastVoteResponse, error) {
	panic("not implemented")
}

func (s DynamodbServer) CreateVoteable(ctx context.Context, r *api.CreateVoteableRequest) (*api.CreateVoteableResponse, error) {
	panic("not implemented")
}

func (s DynamodbServer) ListVoteables(context.Context, *api.ListVoteablesRequest) (*api.ListVoteablesResponse, error) {
	var results []models.Voteable
	err := s.table.Scan().All(&results)
	if err != nil {
		return nil, err
	}

	var voteables []*api.Voteable

	for _, r := range results {
		// Convert answers map to slice
		var answers = make([]string, len(r.Answers))
		answers = append(answers, r.Answers...)

		voteables = append(voteables, &api.Voteable{
			Uuid:     r.UUID,
			Question: r.Question,
			Answers:  answers,
		})
	}

	return &api.ListVoteablesResponse{
		Votables: voteables,
	}, nil

}
