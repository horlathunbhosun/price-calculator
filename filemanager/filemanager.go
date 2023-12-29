package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Readline(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("unable to open file")
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("reading the file content failed")

	}

	file.Close()

	return lines, nil

}
