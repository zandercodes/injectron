package injector

import "github.com/zandercodes/injectron/utils"

// CreateInjector creates a new injector with the following fields set:
func CreateInjector(executable string, jsFiles []string) *Injector {
	freePort := utils.GetFreeRandomPort() // Get a free random port to use for debugging

	// Return a new injector with the following fields set
	return &Injector{
		debugPort:           freePort,   // Set the debug port to the free random port
		executable:          executable, // Set the executable to the executable passed in
		scripts:             jsFiles,    // Set the scripts to the scripts passed in
		waitingBeforeInject: 1000,       // Set the waiting time to 1000ms
		killIfRunning:       false,      // Set kill if running to false
	}
}

// SetCustomPort sets the debug port to the port passed in
func (i *Injector) SetCustomPort(port int) {
	i.debugPort = port // Set the debug port to the port passed in
}

// SetWaitingTime sets the waiting time to the time passed in
func (i *Injector) SetWaitingTime(waitingTime int) {
	i.waitingBeforeInject = waitingTime // Set the waiting time to the time passed in
}

// SetKillIfRunning sets the kill if running to the value passed in
func (i *Injector) SetKillIfRunning(kill bool) {
	i.killIfRunning = kill // Set the kill if running to the value passed in
}
