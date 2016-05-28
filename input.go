package inquire

import (
	"bufio"
	"fmt"
	"strings"
)

type Input struct {
	Default  string
	Message  string
	Name     string
	Validate func(string) error
	When     func(map[string]string) bool
	Filter   func(string) string
}

func (input *Input) Ask(answers map[string]string, reader *bufio.Reader) error {
	if input.When != nil && !input.When(answers) {
		return nil
	}
	answer, err := input.printAndRead(reader)
	if err != nil {
		return err
	}
	if input.Filter != nil {
		answer = input.Filter(answer)
	}

	answers[input.Name] = answer
	return nil
}

func (input *Input) printAndRead(reader *bufio.Reader) (string, error) {
	fmt.Print(input.formattedMessage())
	answer, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	answer = strings.Replace(answer, "\n", "", -1)

	if answer == "" {
		answer = input.Default
	}

	if input.Validate != nil {
		if err := input.Validate(answer); err != nil {
			fmt.Println(err.Error())
			return input.printAndRead(reader)
		}
	}

	return answer, nil
}

func (input *Input) formattedMessage() string {
	message := fmt.Sprintf("%s ", input.Message)
	if input.Default != "" {
		message += fmt.Sprintf("(%s) ", input.Default)
	}
	return message
}
