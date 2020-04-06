package util

import (
	"errors"
	"log"
	"os/exec"
)

func Build(file string) error {
	// check param
	if file == "" {
		return errors.New("file name is empty")
	}
	// go build -buildmode=plugin *.go
	cmd := exec.Command("go", "build", "-buildmode=plugin", file)
	err := cmd.Run()
	if err != nil {
		log.Printf("command exec err. %v\n", err)
		return err
	}
	return nil
}
