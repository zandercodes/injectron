package injector

import (
	"errors"
	"time"
)

// waitingOfApp is a method that waits for the app to start running
func (i *Injector) waitingOfApp() error {
	now := time.Now() // get the current time
	for {
		// check if the timeout has been reached
		if time.Since(now).Seconds() > float64(i.waitingTimeout) {
			return errors.New("timeout waiting for app to start running")
		}

		// wait for one second before continuing the loop
		time.Sleep(time.Second)

		if _, err := i.getSockets(); err != nil {
			continue // if there are no open windows in the app yet (or the app is not running), continue the loop
		}

		break // if the app is running, break the loop
	}
	return nil
}
