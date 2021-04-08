package app

import (
	"fmt"

	"github.com/fatih/color"
)

type appError struct {
	command string
	text    string
}

func (e *appError) Error() string {
	colorFunc := color.New(color.FgRed).SprintFunc()
	errorMessage := "Oops, " + e.command + " command goes wrong"

	return fmt.Sprintf("%s: %s", colorFunc(errorMessage), e.text)
}

func NewAppError(command string, text string) error {
	return &appError{command, text}
}
