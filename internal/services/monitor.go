package services

import (
	"context"
	"fmt"
	"time"

	"uptime-monitor-go/internal/repositories"
	"uptime-monitor-go/pkg/telegram"
)

type Monitor struct {
	Targets   []string
	Repo      *repositories.StatusRepo
	Telegram  *telegram.Client
	Interval  time.Duration
}

func (m *Monitor) Start(ctx context.Context) {
	ticker := time.NewTicker(m.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for _, target := range m.Targets {
				go m.check(ctx, target)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (m *Monitor) check(ctx context.Context, url string) {
	up := PingURL(ctx, url, 5*time.Second)
	newStatus := "DOWN"
	if up {
		newStatus = "UP"
	}

	prevStatus, _ := m.Repo.Get(url)
	if prevStatus != newStatus {
		m.Repo.Set(url, newStatus)
		m.Telegram.Send(fmt.Sprintf("🔔 %s: %s -> %s", url, prevStatus, newStatus))
	}
}
