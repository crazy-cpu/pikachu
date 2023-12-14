package pikachu

import (
	"fmt"
	"time"
)

type Sync struct {
	Response chan []byte
	err      error
	timer    *time.Timer
}

func newSyncAndWaiting(token string) ([]byte, error) {
	sync := Sync{
		Response: make(chan []byte, 1024),
		timer:    time.NewTimer(5 * time.Second),
	}

	defer func() {
		sync.timer.Stop()
		close(sync.Response)
	}()

	tokenCache.Store(token, sync)
	select {
	case <-sync.timer.C:
		return nil, fmt.Errorf("request of model query timeout with token =%s", token)
	case data := <-sync.Response:
		return data, nil
	}
}

func SendAndWaiting(topic string, token string, payload any) ([]byte, error) {
	err := Publish(topic, payload)
	if err != nil {
		return nil, err
	}

	return newSyncAndWaiting(token)

}
