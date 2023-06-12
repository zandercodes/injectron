package injector

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (i *Injector) getSockets() ([]string, error) {
	// check if the app is running by making a GET request to its debug port
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/json/list?t=%d", i.debugPort, time.Now().Unix()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var windows []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&windows); err != nil {
		return nil, err
	}

	var sockets []string
	for _, window := range windows {
		if window["type"].(string) != "page" {
			continue
		}
		if webSocketDebuggerURL, ok := window["webSocketDebuggerUrl"].(string); ok {
			sockets = append(sockets, webSocketDebuggerURL)
		}
	}
	if len(sockets) == 0 {
		return nil, ErrNoSocketsFound
	}
	return sockets, nil
}
