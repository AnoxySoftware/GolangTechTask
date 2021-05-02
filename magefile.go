// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Protoc generates go files from the service protobuf
func Protoc() error {
	return sh.Run(
		"protoc",
		"--go_out=.", "--go_opt=paths=source_relative",
		"--go-grpc_out=.", "--go-grpc_opt=paths=source_relative",
		"api/service.proto",
	)
}
