package inquire

import (
	"bufio"
	"fmt"
	"strings"
)

type Expand struct {
	Default string
	Message string
	Name    string
	When    func(map[string]string) bool
	Choices []Choice
}

func (exp *Expand) Ask(answers map[string]string, reader *bufio.Reader) error {
	if exp.When != nil && !exp.When(answers) {
		return nil
	}
	answer, err := exp.printAndRead(reader, false)
	if err != nil {
		return err
	}

	answers[exp.Name] = answer
	return nil
}

func (exp *Expand) printAndRead(reader *bufio.Reader, long bool) (string, error) {
	if long {
		fmt.Print(exp.longMessage())
	} else {
		fmt.Print(exp.shortMessage())
	}

	answer, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	answer = strings.ToLower(strings.Replace(answer, "\n", "", -1))

	if answer == "" {
		answer = exp.Default
	} else {
		answer = answer[0:1]
	}

	if answer == "h" {
		return exp.printAndRead(reader, true)
	} else if exp.keyValue(answer) == "" {
		fmt.Printf("\nAnswer must be one of %s\n", exp.stringKeys())
		return exp.printAndRead(reader, false)
	}

	return exp.keyValue(answer), nil
}

func (exp *Expand) shortMessage() string {
	return fmt.Sprintf("\n%s: (%s) ", exp.Message, exp.stringKeys())
}

func (exp *Expand) longMessage() string {
	message := ""

	for _, choice := range exp.Choices {
		message += fmt.Sprintf("  %s) ", choice.Key)
		if choice.Key == exp.Default || choice.Value == exp.Default {
			message += fmt.Sprint("(default) ")
		}
		message += fmt.Sprintf("%s\n", choice.Message)
	}
	message += fmt.Sprintf("  h) Help, list all options\n  Answer: ")

	return message
}

func (exp *Expand) stringKeys() string {
	keys := ""
	for _, choice := range exp.Choices {
		if choice.Key == exp.Default || choice.Value == exp.Default {
			keys += strings.ToUpper(choice.Key)
		} else {
			keys += choice.Key
		}
	}
	keys += "h"
	return keys
}

func (exp *Expand) keyValue(key string) string {
	for _, choice := range exp.Choices {
		if choice.Key == key {
			return choice.Value
		}
	}
	return ""
}
