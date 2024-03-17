package parser

import (
	"bufio"
	"bytes"
	"fmt"
)

func splitLines(text []byte) []string {
	var lines []string

	scanner := bufio.NewScanner(bytes.NewReader(text))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func joinLines(lines []string) []byte {
	var buffer bytes.Buffer

	for _, line := range lines {
		fmt.Fprintln(&buffer, line)
	}

	return buffer.Bytes()
}
