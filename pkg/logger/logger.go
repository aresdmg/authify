package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func Green(message string) {
	now := time.Now()

	timestamp := now.Format("2006-01-02 15:04:05 ")

	fmt.Print(timestamp)
	color.New(color.BgGreen, color.FgBlack).Println(message)
}

func Red(message string) {
	now := time.Now()

	timestamp := now.Format("2006-01-02 15:04:05 ")

	fmt.Print(timestamp)
	color.New(color.BgRed, color.FgBlack).Println(message)
}

func Blue(message string) {
	now := time.Now()

	timestamp := now.Format("2006-01-02 15:04:05 ")

	fmt.Print(timestamp)
	color.New(color.BgBlue, color.FgBlack).Println(message)
}

func Yellow(message string) {
	now := time.Now()

	timestamp := now.Format("2006-01-02 15:04:05 ")

	fmt.Print(timestamp)
	color.New(color.BgYellow, color.FgBlack).Println(message)
}

func White(message string)  {
	now := time.Now()

	timestamp := now.Format("2006-01-02 15:04:05 ")

	fmt.Print(timestamp)
	color.New(color.BgHiWhite, color.FgBlack).Println(message)
}