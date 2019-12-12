package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Usage() {
	fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {

	// Open file
	var fileArgs string
	flag.Usage = Usage
	flag.StringVar(&fileArgs, "file", "variables.tf", "a string var")
	flag.Parse()

	if filepath.Ext(strings.TrimSpace(fileArgs)) != ".tf" {
		fmt.Println("Invalid .tf file")
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(fileArgs)

	// if error while openning
	if err != nil {
		fmt.Println(err)
		Usage()
		os.Exit(1)
	}

	// get content as string
	content := string(bytes)

	re, _ := regexp.Compile(`(?m)variable "([a-zA-Z_-]+)"`)
	for _, match := range re.FindAllStringSubmatch(content, -1) {
		fmt.Println(match[1] + " = \"\"\n")
	}
}
