package main

import (
	"fmt"
	"job4j.ru/go-lang-base/internal/base"
	"job4j.ru/go-lang-base/internal/tracker"
)

func main() {
	fmt.Println("Hello World")
	first := 100
	second := 10
	res := base.Add(first, second)

	fmt.Printf("%d + %d = %d\n", first, second, res)

	ui := tracker.UI{
		In:      tracker.ConsoleInput{},
		Out:     tracker.ConsoleOutput{},
		Tracker: tracker.NewTracker(),
	}
	ui.Run()
}
