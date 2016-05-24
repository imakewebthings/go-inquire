package inquire

import (
	"bufio"
	"os"
)

type Question interface {
	Ask(map[string]string, *bufio.Reader) error
}

var reader = bufio.NewReader(os.Stdin)

func Prompt(questions []Question) (map[string]string, error) {
	var err error
	answers := make(map[string]string)

	for _, q := range questions {
		err = q.Ask(answers, reader)
		if err != nil {
			return nil, err
		}
	}

	return answers, nil
}
