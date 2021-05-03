// +build mage

package main

import (
	"encoding/json"
	"flag"
	"os"
	"strings"

	"github.com/magefile/mage/sh"
)

func GrpcurlCast() error {
	return sh.RunV(
		"grpcurl", "-plaintext", "-proto", "api/service.proto", "localhost:3000", "VotingService/CastVote",
	)
}

func GrpcurlCreate() error {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	question := fs.String("question", "question", "question")
	answers := fs.String("answer", "answer1,answer2", "comma seperated answers")

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
		"-d",
		string(data),
		"-plaintext",
		"-proto", "api/service.proto",
		"localhost:3000", "VotingService/CreateVoteable",
	)
}

func GrpcurlList() error {
	return sh.RunV(
		"grpcurl", "-plaintext", "-proto", "api/service.proto", "localhost:3000", "VotingService.ListVoteables",
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
