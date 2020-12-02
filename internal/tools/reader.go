package tools

import (
	"bufio"
	"os"
)

// Data represents the content of a data file
type Data struct {
	Lines []string
}

// GetData reads a file and returns the content as a list of strings
func GetData(path string) (*Data, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data Data
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data.Lines = append(data.Lines, scanner.Text())
	}
	return &data, scanner.Err()
}
