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

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("Could not open file")
	}

	defer file.Close() // execute once the function executed successfully (or returns an error)

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		// file.Close()
		return nil, errors.New("Reading the content file failed")
	}

	// file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	defer file.Close() // execute once the function executed successfully (or returns an error)

	if err != nil {
		return errors.New("Could not create file")
	}

	time.Sleep(3 * time.Second) // simulate a long process

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data) // encode data to JSON and write to file

	if err != nil {
		// file.Close()
		return errors.New("Failed to convert data to JSON")
	}

	// file.Close()
	return nil
}

func New(inputFilePath, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}
