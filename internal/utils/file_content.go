package utils

import (
	"bytes"
	"errors"
	"os"
)

func RepeatFileContent(path string, count int) ([]string, error) {
	if count < 0 {
		return nil, errors.New("count must be a non negative integer")
	}

	read := func() (string, error) {
		file, err := os.Open(path)

		if err != nil {
			return "", err
		}

		defer file.Close()

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(file)

		if err != nil {
			return "", err
		}

		return buf.String(), nil
	}

	result := make([]string, 0, count)

	for i := 0; i < count; i++ {
		str, err := read()

		if err != nil {
			return nil, err
		}

		result = append(result, str)
	}

	return result, nil
}
