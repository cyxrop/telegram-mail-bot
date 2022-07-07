package worker

import (
	"context"
	"log"
	"time"

	service "gitlab.ozon.dev/cyxrop/homework-2/internal/app/service/user"
)

type Notifier struct {
	us service.UserService

	every time.Duration
}

func NewNotifier(us service.UserService, every time.Duration) *Notifier {
	return &Notifier{
		us:    us,
		every: every,
	}
}

func (n *Notifier) Run(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("notifier ctx done")
				return
			case <-time.After(n.every):
				log.Println("Notify all users")
				if errs := n.us.NotifyAll(ctx); len(errs) != 0 {
					log.Printf("notify all: %s\n", errs)
				}
			}
		}
	}()
}
