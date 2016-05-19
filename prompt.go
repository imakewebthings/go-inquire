package inquire

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type Input struct {
	Default  string
	Message  string
	Name     string
	Validate func(string) error
	When     func(map[string]string) bool
	Filter   func(string) string
}

func Prompt(questions []Input) (map[string]string, error) {
	answers := make(map[string]string)

	for _, q := range questions {
		if q.When != nil && !q.When(answers) {
			continue
		}
		answer, err := ask(&q)
		if err != nil {
			return nil, err
		}
		if q.Filter != nil {
			answer = q.Filter(answer)
		}
		answers[q.Name] = answer
	}

	return answers, nil
}

func ask(q *Input) (string, error) {
	fmt.Printf(q.formattedMessage())
	answer, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	answer = strings.Replace(answer, "\n", "", -1)

	if answer == "" {
		answer = q.Default
	}

	if q.Validate != nil {
		if err := q.Validate(answer); err != nil {
			fmt.Println(err.Error())
			return ask(q)
		}
	}

	return answer, nil
}

func (i Input) formattedMessage() string {
	message := fmt.Sprintf("%s ", i.Message)
	if i.Default != "" {
		message = message + fmt.Sprintf("(%s) ", i.Default)
	}
	return message
}
