package main

import (
	"fmt"
	"job4j.ru/go-lang-base/internal/base"
)

func main() {
	fmt.Println("Hello World")
	first := 100
	second := 10
	res := base.Add(first, second)
	fmt.Println(fmt.Sprintf("%d + %d = %d", first, second, res))
}
