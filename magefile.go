// +build mage

package main

import (
	"encoding/json"
	"flag"
	"os"
	"strings"

	"github.com/magefile/mage/sh"
	"github.com/peterbourgon/ff/v3"
)

// GrpcurlCast calls the VotingService CastVote endpoint
func GrpcurlCast() error {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	index := fs.Int64("index", 1, "answer index")
	uuid := fs.String("uuid", "", "uuid")

	data, err := json.Marshal(struct {
		UUID  string `json:"uuid"`
		Index int64  `json:"answer_index"`
	}{
		*uuid,
		*index,
	})
	if err != nil {
		return err
	}

	return sh.RunV(
		"grpcurl",
		"-d",
		string(data),
		"-plaintext",
		"-proto", "api/service.proto",
		"localhost:3000", "VotingService/CastVote",
	)
}

// GrpcurlCreate calls the VotingService CreateVoteable endpoint
func GrpcurlCreate() error {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	question := fs.String("question", "question", "question")
	answers := fs.String("answer", "answer1,answer2", "comma seperated answers")
	ff.Parse(fs, os.Args[1:], ff.WithEnvVarNoPrefix())

	data, err := json.Marshal(struct {
		Question string   `json:"question"`
		Answers  []string `json:"answers"`
	}{
		*question,
		strings.Split(*answers, ","),
	})
	if err != nil {
		return err
	}

	return sh.RunV(
		"grpcurl",
		"-d", string(data),
		"-plaintext",
		"-proto", "api/service.proto",
		"localhost:3000", "VotingService/CreateVoteable",
	)
}

// GrpcurlList calls the VotingService ListVoteables endpoint
func GrpcurlList() error {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	pageToken := fs.String("pagetoken", "", "Page Token")
	pageSize := fs.Int64("pagesize", 10, "Page Size")
	ff.Parse(fs, os.Args[1:], ff.WithEnvVarNoPrefix())

	data, err := json.Marshal(struct {
		PageToken string `json:"page_token"`
		PageSize  int64  `json:"page_size"`
	}{
		*pageToken,
		*pageSize,
	})
	if err != nil {
		return err
	}

	return sh.RunV(
		"grpcurl",
		"-d", string(data),
		"-plaintext",
		"-proto", "api/service.proto",
		"localhost:3000", "VotingService.ListVoteables",
	)
}

// Protoc generates go files from the service protobuf
func Protoc() error {
	return sh.Run(
		"protoc",
		"--go_out=.", "--go_opt=paths=source_relative",
		"--go-grpc_out=.", "--go-grpc_opt=paths=source_relative",
		"api/service.proto",
	)
}

// Test runs go test ./...
func Test() error {
	return sh.Run("go", "test", "./...")
}
