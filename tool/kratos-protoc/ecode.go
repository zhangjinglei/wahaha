package main

import (
	"os/exec"
)

const (
	_getEcodeGen = "go get -u github.com/zhangjinglei/wahaha/tool/protobuf/protoc-gen-ecode@" + Version
	_ecodeProtoc = "protoc --proto_path=%s --proto_path=%s --proto_path=%s --ecode_out=:."
)

func installEcodeGen() error {
	if _, err := exec.LookPath("protoc-gen-ecode"); err != nil {
		if err := goget(_getEcodeGen); err != nil {
			return err
		}
	}
	return nil
}

func genEcode(files []string) error {
	return generate(_ecodeProtoc, files)
}
