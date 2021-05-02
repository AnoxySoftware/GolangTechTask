package main

import (
	"flag"
	"net"
	"os"
	"os/signal"

	pb "github.com/Sacro/GolangTechTask/api"
	"github.com/Sacro/GolangTechTask/internal/server"
	"github.com/apex/log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/peterbourgon/ff/v3"
	"google.golang.org/grpc"
)

func main() {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var (
		address  = fs.String("address", ":3000", "gRPC address")
		endpoint = fs.String("endpoint", "http://127.0.0.1:8000", "DynamoDB endpoint")
		id       = fs.String("id", "id", "AWS account ID")
		region   = fs.String("region", "region", "AWS Region")
		secret   = fs.String("secret", "secret", "AWS secret")
		token    = fs.String("token", "token", "AWS token")
	)

	err := ff.Parse(fs, os.Args[1:])
	if err != nil {
		log.WithError(err).Fatal("parsing flags")
	}
	session, err := session.NewSession()
	if err != nil {
		log.WithError(err).Fatal("creating session")
	}

	s := grpc.NewServer()

	config := aws.NewConfig().WithEndpoint(*endpoint).WithRegion(*region).WithCredentials(credentials.NewStaticCredentials(*id, *secret, *token))
	db := dynamo.New(session, config)
	srv := server.NewDynamodbServer(db)

	pb.RegisterVotingServiceServer(s, srv)

	listener, err := net.Listen("tcp", *address)
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
