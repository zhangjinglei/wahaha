package main

import (
	"os/exec"
)

const (
	_getGoGen = "go get -u google.golang.org/protobuf/cmd/protoc-gen-go"
)

func installGenGo() error {
	if _, err := exec.LookPath("protoc-gen-go"); err != nil {
		if err := goget(_getGoGen); err != nil {
			return err
		}
	}
	return nil
}
