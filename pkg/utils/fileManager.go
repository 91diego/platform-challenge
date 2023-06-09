package utils

import (
	"bufio"
	"os"
)

// FileReader
func FileReader(file *string) ([]string, error) {
	readFile, err := os.Open(*file)
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	return lines, nil
}
