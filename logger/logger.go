package logger

import (
	"fmt"

	"github.com/fatih/color"
)

func Error(msg string, any ...any) {
	color.New(color.Bold, color.FgRed).Println(fmt.Sprintf(msg, any...))
}

func Info(msg string, any ...any) {
	color.New(color.Bold, color.FgBlue).Println(fmt.Sprintf(msg, any...))
}

func Warning(msg string, any ...any) {
	color.New(color.Bold, color.FgYellow).Println(fmt.Sprintf(msg, any...))
}

func Subtle(msg string, any ...any) {
	color.New(color.Faint, color.Italic).Println(fmt.Sprintf(msg, any...))
}
