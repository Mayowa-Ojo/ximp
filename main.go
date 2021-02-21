package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	var pathToFile string
	args := os.Args[1:]

	if len(args) < 1 {
		pathToFile = ".env.development"
	} else {
		pathToFile = args[0]
	}

	file, err := os.Open(pathToFile)

	if err != nil {
		log.Fatal("[Error]: failed to open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		key := strings.Split(line, "=")[0]
		lines = append(lines, key)
	}

	var fileDest string

	if len(args) < 1 {
		fileDest = ".env.example"
	} else {
		path := strings.Split(pathToFile, "/")
		dir := strings.Join(path[:len(path)-1], "/")
		fileDest = dir + "/.env.example"
	}

	f, err := os.Create(fileDest)
	if err != nil {
		log.Fatal("[Error]: couldn't create file")
	}

	defer f.Close()

	for index, line := range lines {
		lines[index] = line + "="
	}

	content := strings.Join(lines, "\n")
	_, err = f.WriteString(content)
	if err != nil {
		log.Fatal("[Error]: couldn't write to file")
	}
}
