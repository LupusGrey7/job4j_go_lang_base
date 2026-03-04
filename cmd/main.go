package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"job4j.ru/go-lang-base/internal/api"
	"job4j.ru/go-lang-base/internal/base"
	"job4j.ru/go-lang-base/internal/config"
	"job4j.ru/go-lang-base/internal/db"
	"job4j.ru/go-lang-base/internal/repository"
	"log"
)

const (
	APIPrefix = "/api"
)

func main() {
	fmt.Println("Hello World")
	first := 100
	second := 10
	res := base.Add(first, second)

	fmt.Printf("%d + %d = %d\n", first, second, res)

	// ---Tracker UI---
	//DB Config
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

	//Server
	server := api.NewServer(repo)

	app := fiber.New() // создаём новый web-сервер
	server.Route(app.Group(APIPrefix))

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
