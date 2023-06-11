package injector

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// waitingOfApp is a method that waits for the app to start running
func (i *Injector) waitingOfApp() {
	isStarted := false // initialize a variable to keep track of whether the app has started running
	for {
		// wait for one second before continuing the loop
		time.Sleep(time.Second)

		// check if the app is running by making a GET request to its debug port
		resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/json/list?t=%d", i.debugPort, time.Now().Unix()))
		if err != nil {
			continue // if the app is not running, continue the loop
		}
		defer resp.Body.Close()

		var windows []map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&windows); err != nil {
			continue // if there are no open windows in the app yet (or the app is not running), continue the loop
		}

		for _, window := range windows {
			if window["type"].(string) != "page" {
				continue // if the window is not a page, continue the loop
			}
			isStarted = true // if we've found a page window, the app has started running
			break
		}

		if isStarted {
			break // if the app is running, break out of the loop
		}
	}
}
