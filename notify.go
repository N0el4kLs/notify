package notify

import (
	"sync"

	"notify/pkg/notifications"
)

type Sender notifications.Notifier

func Notice(providers ...Sender) {
	var wg sync.WaitGroup
	for _, provider := range providers {
		wg.Add(1)
		go func(p Sender) {
			defer wg.Done()
			p.Notice()
		}(provider)
	}
	wg.Wait()
}
