package notify

import (
	"sync"

	"notify/pkg/notifications"
)

type Provider notifications.Notifier

func Notice(providers ...Provider) {
	var wg sync.WaitGroup
	for _, provider := range providers {
		wg.Add(1)
		go func(p Provider) {
			defer wg.Done()
			p.Notice()
		}(provider)
	}
	wg.Wait()
}
