package injector

// Injector is the main struct for the injector
type Injector struct {
	debugPort           int      // debugPort is the port that the injector will listen on
	executable          string   // executable is the application that will be injected
	scripts             []string // scripts is the list of scripts that will be injected
	waitingBeforeInject int      // waitingBeforeInject is the time in milliseconds that the injector will wait before injecting
	killIfRunning       bool     // killIfRunning is a flag that will kill the application if it is already running
}
