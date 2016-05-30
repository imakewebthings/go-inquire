package inquire

import (
	"bufio"
	"fmt"
	"strings"
)

type Confirm struct {
	Default string
	Message string
	Name    string
	When    func(map[string]string) bool
}

var validAnswers map[string]bool = map[string]bool{
	"yes": true,
	"y":   true,
	"no":  true,
	"n":   true,
}

func (c *Confirm) Ask(answers map[string]string, reader *bufio.Reader) error {
	if c.When != nil && !c.When(answers) {
		return nil
	}
	answer, err := c.printAndRead(reader)
	if err != nil {
		return err
	}

	answers[c.Name] = answer
	return nil
}

func (c *Confirm) printAndRead(reader *bufio.Reader) (string, error) {
	fmt.Print(c.formattedMessage())

	answer, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	answer = strings.ToLower(strings.Replace(answer, "\n", "", -1))

	if answer == "" {
		answer = c.Default
	}

	if !validAnswers[answer] {
		fmt.Println("Answer must by y/n")
		return c.printAndRead(reader)
	}

	return answer, nil
}

func (c *Confirm) formattedMessage() string {
	y := "y"
	n := "n"
	if c.Default == "y" || c.Default == "yes" {
		y = "Y"
	} else if c.Default == "n" || c.Default == "no" {
		n = "N"
	}

	return fmt.Sprintf("%s (%s/%s) ", c.Message, y, n)
}
