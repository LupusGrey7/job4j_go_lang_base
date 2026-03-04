package main

import (
	"context"
	"fmt"
	"job4j.ru/go-lang-base/internal/base"
	"job4j.ru/go-lang-base/internal/config"
	"job4j.ru/go-lang-base/internal/db"
	"job4j.ru/go-lang-base/internal/repository"
	"job4j.ru/go-lang-base/internal/tracker"
	"log"
)

func main() {
	fmt.Println("Hello World")
	first := 100
	second := 10
	res := base.Add(first, second)

	fmt.Printf("%d + %d = %d\n", first, second, res)

	//Tracker UI
	ctx := context.Background()

	cfg := db.Config{
		Host:     config.Env("DB_HOST", "localhost"),
		Port:     config.EnvInt("DB_PORT", 6543),
		User:     config.Env("DB_USER", "postgres"),
		Password: config.Env("DB_PASSWORD", "password"),
		DBName:   config.Env("DB_NAME", "tracker"),
		SSLMode:  config.Env("DB_SSLMODE", "disable"),
	}

	pool, err := db.NewPool(ctx, cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repo := repository.NewRepoPg(pool)

	ui := tracker.UI{
		In:    tracker.ConsoleInput{},
		Out:   tracker.ConsoleOutput{},
		Store: repo,
	}

	if err := ui.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
