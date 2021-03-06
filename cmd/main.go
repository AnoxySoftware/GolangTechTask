package main

import (
	"context"
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
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var (
		address  = fs.String("address", ":3000", "gRPC address")
		endpoint = fs.String("endpoint", "http://127.0.0.1:8000", "DynamoDB endpoint")
		id       = fs.String("id", "id", "AWS account ID")
		level    = fs.String("level", log.InfoLevel.String(), "Log level")
		region   = fs.String("region", "region", "AWS Region")
		secret   = fs.String("secret", "secret", "AWS secret")
		token    = fs.String("token", "token", "AWS token")
	)

	err := ff.Parse(fs, os.Args[1:], ff.WithEnvVarNoPrefix())
	if err != nil {
		log.WithError(err).Fatal("parsing flags")
	}

	if l, e := log.ParseLevel(*level); e != nil {
		log.WithError(e).Fatal("setting log level")
	} else {
		log.SetLevel(l)
		log.WithField("level", *level).Info("set log level")
	}

	session, err := session.NewSession()
	if err != nil {
		log.WithError(err).Fatal("creating session")
	}
	log.Info("created session")

	tp := sdktrace.NewTracerProvider()
	defer func() { _ = tp.Shutdown(ctx) }()

	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor(otelgrpc.WithTracerProvider(tp))),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor(otelgrpc.WithTracerProvider(tp))),
	)

	config := aws.NewConfig().WithEndpoint(*endpoint).WithRegion(*region).WithCredentials(credentials.NewStaticCredentials(*id, *secret, *token))
	log.WithField("config", config).Debug("dynamodb config")
	db := dynamo.New(session, config)
	srv := server.NewDynamodbServer(ctx, db)

	pb.RegisterVotingServiceServer(s, srv)
	log.Info("registered voting service server")

	listener, err := net.Listen("tcp", *address)
	if err != nil {
		log.WithError(err).Fatal("establishing listener")
	}
	log.Info("established listener")

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

	select {
	case err := <-errChan:
		log.WithError(err).Fatal("running server")
	case <-stopChan:
		log.Info("cancelling via context")
		cancel()
		log.Info("graceful stop in progress")
		s.GracefulStop()
	}
}
