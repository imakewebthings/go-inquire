package inquire

import (
	"bufio"
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

type Password struct {
	Message  string
	Name     string
	Validate func(string) error
	When     func(map[string]string) bool
}

func (pw *Password) Ask(answers map[string]string, reader *bufio.Reader) error {
	if pw.When != nil && !pw.When(answers) {
		return nil
	}
	answer, err := pw.printAndRead(reader)
	if err != nil {
		return err
	}

	answers[pw.Name] = answer
	return nil
}

func (pw *Password) printAndRead(reader *bufio.Reader) (string, error) {
	fmt.Print(pw.formattedMessage())
	byteAnswer, err := terminal.ReadPassword(syscall.Stdin)

	if err != nil {
		fmt.Println(err.Error())
		return pw.printAndRead(reader)
	}
	answer := string(byteAnswer)

	if pw.Validate != nil {
		if err := pw.Validate(answer); err != nil {
			fmt.Println(err.Error())
			return pw.printAndRead(reader)
		}
	}

	return answer, nil
}

func (pw *Password) formattedMessage() string {
	return fmt.Sprintf("%s ", pw.Message)
}
