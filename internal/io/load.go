package io

import (
	"bufio"
	"errors"
	"os"
)

func LoadKeys(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var keys []string
	scanner := bufio.NewScanner(f)
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		if lineNo == 1 { // skip ruler
			continue
		}
		s := scanner.Text()
		if len(s) != 16 {
			return nil, errors.New("line must be exactly 16 chars")
		}
		// reject non-ASCII
		for i := 0; i < 16; i++ {
			if s[i] > 127 {
				return nil, errors.New("non-ASCII character")
			}
		}
		keys = append(keys, s) // keep exact 16 bytes (padded with spaces)
	}
	return keys, scanner.Err()
}
