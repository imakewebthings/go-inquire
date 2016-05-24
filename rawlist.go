package inquire

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type RawList struct {
	Default string
	Message string
	Name    string
	When    func(map[string]string) bool
	Filter  func(string) string
	Choices []Choice
}

func (list *RawList) Ask(answers map[string]string, reader *bufio.Reader) error {
	if list.When != nil && !list.When(answers) {
		return nil
	}
	answer, err := list.printAndRead(reader)
	if err != nil {
		return err
	}
	if list.Filter != nil {
		answer = list.Filter(answer)
	}

	answers[list.Name] = answer
	return nil
}

func (list *RawList) printAndRead(reader *bufio.Reader) (string, error) {
	fmt.Printf(list.formattedMessage())

	strIndex, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return list.printAndRead(reader)
	}

	strIndex = strings.TrimSpace(strIndex)
	index, err := strconv.Atoi(strIndex)
	if err != nil || index < 1 || index > len(list.Choices) {
		fmt.Printf("\nPlease pick a number between 1 and %d.\n", len(list.Choices))
		return list.printAndRead(reader)
	}

	return list.Choices[index-1].Value, nil
}

func (list *RawList) formattedMessage() string {
	message := fmt.Sprintf("\n%s\n", list.Message)
	for i, choice := range list.Choices {
		message += fmt.Sprintf("%d) %s\n", i+1, choice.Message)
	}
	message += fmt.Sprint("Answer: ")
	return message
}
