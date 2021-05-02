package server

import (
	"context"
	"net"
	"testing"

	"github.com/apex/log"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/Sacro/GolangTechTask/api"
)

const (
	bufSize = 1024 * 1024
)

// Verify our service implements the required methods
var _ api.VotingServiceServer = MockServer{}

type MockServerTestSuite struct {
	suite.Suite

	client api.VotingServiceClient
}

func (suite *MockServerTestSuite) SetupTest() {
	log.SetLevel(log.InfoLevel)

	listener := bufconn.Listen(bufSize)
	s := grpc.NewServer()
	api.RegisterVotingServiceServer(s, &MockServer{})
	go func() {
		if err := s.Serve(listener); err != nil {
			log.WithError(err).Fatalf("Server exited with error: %v", err)
		}
	}()

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}), grpc.WithInsecure())
	suite.Nil(err)

	suite.client = api.NewVotingServiceClient(conn)
}

func (suite *MockServerTestSuite) TestListVoteables() {
	ctx := context.Background()

	question := "What is your favourite color"
	answers := []string{"Red", "Blue", "Green"}

	createRes, err := suite.client.CreateVoteable(ctx, &api.CreateVoteableRequest{
		Question: question,
		Answers:  answers,
	})
	if !suite.Nil(err) {
		log.WithError(err).Error("creating voteable")
		return
	}

	suite.NotEmpty(createRes.Uuid)

	listRes, err := suite.client.ListVoteables(ctx, &api.ListVoteablesRequest{})
	if !suite.Nil(err) {
		log.WithError(err).Error("listing voteables")
		return
	}

	if !suite.Len(listRes.Votables, 1) {
		return
	}

	v := listRes.Votables[0]
	suite.Equal(question, v.Question)
	suite.Equal(answers, v.Answers)
	suite.Equal(createRes.Uuid, v.Uuid)
}

func TestMockServerTestSuite(t *testing.T) {
	suite.Run(t, new(MockServerTestSuite))
}
