package contract

import (
	"errors"
	"log"
	"os"
	"path"
	"plugin"
	"strings"
	"xchain-contract-go/src/xchain/config"
	"xchain-contract-go/src/xchain/util"
)

var library *plugin.Plugin

func Build(content string) (string, error) {
	if content == "" {
		msg := "content is empty"
		return msg, errors.New(msg)
	}
	// check or delete library
	if existsLibrary(config.WorkDir) {
		deleteLibrary(config.WorkDir)
	}
	// build to library
	if !buildLibrary(config.WorkDir, content) {
		log.Printf("FAILED TO BUILD LIBRARY\n")
		msg := "wrong when build library file"
		return msg, errors.New(msg)
	}
	log.Printf("SUCCEEDED TO BUILD LIBRARY, NOW YOU CAN INVOKE...\n")
	return "", nil
}

func Load() (string, error) {
	// check library
	if !existsLibrary(config.WorkDir) {
		msg := "library file does not exist"
		return msg, errors.New(msg)
	}
	// load to library
	if !loadLibrary(config.WorkDir) {
		// exit
		log.Fatalf("Wrong when loading library file")
	}
	return "", nil
}

// build library file (*.so)
func buildLibrary(dir, content string) bool {
	if dir == "" || content == "" {
		return false
	}
	// replace package name
	exp := `package \w*`
	rep := "package main"
	content, err := util.Replace(content, exp, rep)
	if err != nil {
		log.Printf("replace package name err. %v\n", err)
		return false
	}
	// write to file
	err = util.Write(config.ContractFileName, content)
	if err != nil {
		log.Printf("write library file err. %v\n", err)
		return false
	}
	// build to *.so file
	err = util.Build(config.ContractFileName)
	if err != nil {
		log.Printf("build library file err. %v\n", err)
		return false
	}
	library, err = plugin.Open(config.LibraryFileName)
	if err != nil {
		log.Printf("open library file err. %v\n", err)
		return false
	}
	return true
}

// load library file (*.so)
func loadLibrary(dir string) bool {
	if dir == "" {
		return false
	}
	// load to *.so file
	var err error
	library, err = plugin.Open(config.LibraryFileName)
	if err != nil {
		log.Printf("open library file err. %v\n", err)
		return false
	}
	return true
}

// check library files (*.so)
func existsLibrary(dir string) bool {
	if dir == "" {
		return false
	}
	// filter *.so file
	directory, err := os.Open(dir)
	if err != nil {
		log.Printf("open library directory err. %v\n", err)
		return false
	}
	list, err := directory.Readdir(-1)
	if err != nil {
		log.Printf("list library directory err. %v\n", err)
		return false
	}
	if len(list) == 0 {
		return false
	}
	for _, v := range list {
		if v == nil {
			continue
		}
		if v.IsDir() {
			continue
		}
		if strings.HasSuffix(v.Name(), ".so") {
			return true
		}
	}
	return false
}

// delete library files (*.so)
func deleteLibrary(dir string) bool {
	if dir == "" {
		return true
	}
	// delete *.so file
	directory, err := os.Open(dir)
	if err != nil {
		log.Printf("open library directory err. %v\n", err)
		return false
	}
	list, err := directory.Readdir(-1)
	if err != nil {
		log.Printf("list library directory err. %v\n", err)
		return false
	}
	if len(list) == 0 {
		return true
	}
	for _, v := range list {
		if v == nil {
			continue
		}
		if v.IsDir() {
			continue
		}
		if !strings.HasSuffix(v.Name(), ".so") {
			continue
		}
		file := path.Join(dir, v.Name())
		err := os.Remove(file)
		if err != nil {
			log.Printf("remove library %s err. %v\n", file, err)
		}
	}
	return true
}
