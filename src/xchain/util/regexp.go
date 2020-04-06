package util

import (
	"errors"
	"log"
	"regexp"
)

func Replace(content, exp, rep string) (string, error) {
	// check params
	if content == "" {
		return "", errors.New("content is empty")
	}
	if exp == "" {
		return "", errors.New("regular expression is empty")
	}
	if rep == "" {
		return "", errors.New("replacement is empty")
	}

	// replace
	re, err := regexp.Compile(exp)
	if err != nil {
		log.Printf("regexp compile err. %v\n", err)
		return "", err
	}
	out := re.ReplaceAllString(content, rep)
	return out, nil
}
