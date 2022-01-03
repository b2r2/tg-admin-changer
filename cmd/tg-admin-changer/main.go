package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/b2r2/tg-admin-changer/internal/repositories"

	"github.com/b2r2/tg-admin-changer/internal/app"
	"github.com/b2r2/tg-admin-changer/internal/config"
)

func main() {
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

	if err := config.Load(); err != nil {
		log.Fatalln(err)
	}

	cfg := config.Get()

	r := repositories.NewRepository(cfg.GetDBConnection(), cfg.GetRedis())

	bot, err := app.New(cfg.GetLogger(), cfg.GetToken(), r)
	if err != nil {
		log.Fatalln(err)
	}

	go bot.Start(ctx)

	<-ctx.Done()

	if err = r.Close(); err != nil {
		log.Fatalln(err)
	}

	bot.Stop()
}
