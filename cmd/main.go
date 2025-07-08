package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"uptime-monitor-go/internal/utils"
	"uptime-monitor-go/internal/repositories"
	"uptime-monitor-go/internal/services"
	"uptime-monitor-go/internal/handlers"
	"uptime-monitor-go/pkg/telegram"
)

func main() {
	cfg, err := utils.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatal("Config load error: ", err)
	}

	repo := repositories.NewStatusRepo()
	tg := telegram.NewClient(cfg.TelegramToken, cfg.TelegramChatID)
	handlers.InitMetrics()

	targets := make([]string, len(cfg.Targets))
	for i, t := range cfg.Targets {
		targets[i] = t.URL
	}

	monitor := services.Monitor{
		Targets:  targets,
		Repo:     repo,
		Telegram: tg,
		Interval: time.Duration(cfg.IntervalSeconds) * time.Second,
	}

	ctx := context.Background()
	go monitor.Start(ctx)

	http.Handle("/metrics", handlers.MetricsHandler())
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
