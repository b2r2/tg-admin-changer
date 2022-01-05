package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/b2r2/tg-admin-changer/internal/app"
	"github.com/b2r2/tg-admin-changer/internal/config"
)

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	go func() {
		select {
		case <-ctx.Done():
			return
		case <-sigCh:
			cancel()
		}
	}()

	cfg := config.Get()

	bot, err := app.New(cfg.GetLogger(), cfg.GetToken())
	if err != nil {
		log.Fatalln(err)
	}

	go bot.Run()

	<-ctx.Done()

	bot.Stop()
	return nil
}

func main() {
	if err := config.Load(); err != nil {
		log.Fatalln(err)
	}

	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
