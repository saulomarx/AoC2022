package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadLines(path string) []string {
	var input []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Deu ruim")
	}
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)

		input = append(input, line)
	}

	return input
}

func ReadLinesNoTrim(path string) []string {
	var input []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Deu ruim")
	}
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = line[:len(line)-1]
		input = append(input, line)
	}

	return input
}
