package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) Readline() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		// fmt.Println(err)
		return nil, errors.New("unable to open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("reading the file content failed")

	}

	return lines, nil

}

func (fm FileManager) WriteResult(data interface{}) error {
	fileVal, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("unable to create file")
	}

	defer fileVal.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(fileVal)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("unable to create json file")
	}
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
