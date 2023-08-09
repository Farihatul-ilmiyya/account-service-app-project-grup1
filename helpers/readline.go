package helpers

import (
	"bufio"
	"os"
)

func Readline() (string, error) {
	var str string
	var err error
	reader := bufio.NewReader(os.Stdin)
	str, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// str = strings.TrimSuffix(str, "\n")
	return str, nil
}
