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
