package util

import (
	"errors"
	"log"
	"os"
	"strings"
)

func Write(file, content string) error {
	// check params
	if file == "" || strings.Contains(file, "/") {
		return errors.New("file name is wrong")
	}
	if content == "" {
		return errors.New("file content is empty")
	}
	// create file
	f, err := os.Create(file)
	defer f.Close()
	if err != nil {
		log.Printf("create file err. %v\n", err)
		return err
	}
	// write to file
	_, err = f.WriteString(content)
	if err != nil {
		log.Printf("write string to file err. %v\n", err)
		return err
	}
	return nil
}
