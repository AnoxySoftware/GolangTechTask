package server

import (
	"context"

	"github.com/Sacro/GolangTechTask/api"
	"github.com/Sacro/GolangTechTask/internal/models"
	"github.com/apex/log"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
)

const tableName = "voteables"

type DynamodbServer struct {
	api.UnimplementedVotingServiceServer

	db    *dynamo.DB
	table dynamo.Table
}

func NewDynamodbServer(ctx context.Context, db *dynamo.DB) *DynamodbServer {
	err := db.CreateTable(tableName, models.Voteable{}).RunWithContext(ctx)
	if err != nil {
		// If ResourceInUseException is thrown our table already exists
		log.WithField("err", err.(awserr.Error).Code()).Info("arg")
		if err.(awserr.Error).Code() == "ResourceInUseException" {
			log.Info("table exists")
		} else {
			log.WithError(err).Fatal("create table")
		}
	} else {
		log.WithError(err).Info("created table")
	}

	table := db.Table(tableName)

	return &DynamodbServer{
		db:    db,
		table: table,
	}
}

func (s DynamodbServer) CastVote(ctx context.Context, r *api.CastVoteRequest) (*api.CastVoteResponse, error) {

	err := s.table.Update("UUID", r.Uuid).SetExpr("votes.#answerIndex", "votes + :val").RunWithContext(ctx)
	if err != nil {
		return nil, err
	}

	return &api.CastVoteResponse{}, nil
}

func (s DynamodbServer) CreateVoteable(ctx context.Context, r *api.CreateVoteableRequest) (*api.CreateVoteableResponse, error) {
	u, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	err = s.table.Put(models.Voteable{
		UUID:     u.String(),
		Votes:    make(map[string]int64),
		Question: r.Question,
		Answers:  r.Answers,
	}).RunWithContext(ctx)

	if err != nil {
		return nil, err
	}

	return &api.CreateVoteableResponse{
		Uuid: u.String(),
	}, nil
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
