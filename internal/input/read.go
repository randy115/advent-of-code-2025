package input

import (
	"bufio"
	"os"
)

func ReadFile(path string) ([]string, error) {
	data, err := os.Open("inputs/" + path)
	if err != nil {
		return nil, err
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func Create2DMatrix(path string) ([][]byte, error) {
	data, err := os.Open("inputs/" + path)
	if err != nil {
		return nil, err
	}

	defer data.Close()
	var grid [][]byte
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}
