package daemon

import (
	"log"
	"sync"
)

// New builds a new daemon
func New() *Daemon {
	return &Daemon{
		bootProcesses: make([]Process, 0),
		mainProcesses: make([]Process, 0),
	}
}

// Daemon is a long running process that will trigger other processes when it starts
type Daemon struct {
	bootProcesses []Process
	mainProcesses []Process
}

// Process is a process that can be set to run at a point in the point process
type Process struct {
	Name     string
	Disabled bool
	Action   func() error
}

// Call logs the process name and calls the action
func (process Process) Call() error {
	if process.Disabled {
		log.Printf("Disabled %v", process.Name)
		return nil
	}

	log.Printf("Starting Daemon Process: %v", process.Name)

	err := process.Action()
	if err != nil {
		return err
	}

	log.Printf("Finished Daemon Process: %v", process.Name)

	return nil
}

// OnStart is a process that will run on start up, these processes are blocking
func (daemon *Daemon) OnStart(process Process) *Daemon {
	daemon.bootProcesses = append(daemon.bootProcesses, process)

	return daemon
}

// Do is a process that will run by the deamon after the boot processes in its own goroutine
func (daemon *Daemon) Do(process Process) *Daemon {
	daemon.mainProcesses = append(daemon.mainProcesses, process)

	return daemon
}

// Start starts the daemon.
/*
	First runs the OnStart processes in serial as blocking.
	Then runs the Do processes in parallel.
*/
func (daemon Daemon) Start() {
	for _, bootProcess := range daemon.bootProcesses {
		err := bootProcess.Call()
		if err != nil {
			panic(err)
		}
	}

	for _, process := range daemon.mainProcesses {
		go process.Call()
	}

	for {
		wg := sync.WaitGroup{}
		wg.Add(1)
		wg.Wait()
	}
}
