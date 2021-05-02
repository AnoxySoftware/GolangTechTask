package main

import (
	"net"
	"os"
	"os/signal"

	"github.com/Sacro/GolangTechTask/internal/server"
	"github.com/apex/log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"google.golang.org/grpc"

	pb "github.com/Sacro/GolangTechTask/api"
)

const (
	endpoint = "http://127.0.0.1:8000"
	region   = "region"
	id       = "id"
	secret   = "secret"
	token    = "token"
	port     = ":3000"
)

func main() {
	session, err := session.NewSession()
	if err != nil {
		log.WithError(err).Fatal("creating session")
	}

	s := grpc.NewServer()

	config := aws.NewConfig().WithEndpoint(endpoint).WithRegion(region).WithCredentials(credentials.NewStaticCredentials(id, secret, token))
	db := dynamo.New(session, config)
	srv := server.NewDynamodbServer(db)

	pb.RegisterVotingServiceServer(s, srv)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.WithError(err).Fatal("establishing listener")
	}

	stopChan := make(chan os.Signal, 1)
	errChan := make(chan error)

	// bind OS events to signal channel
	signal.Notify(stopChan, os.Interrupt)

	// run blocking call in a separate goroutine, report errors via channel
	go func() {
		if err := s.Serve(listener); err != nil {
			errChan <- err
		}
	}()

	defer func() {
		s.GracefulStop()
	}()

	select {
	case err := <-errChan:
		log.WithError(err).Fatal("running server")
	case <-stopChan:
		log.Info("stopping server")
	}
}
