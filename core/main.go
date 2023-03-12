package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/KarnerTh/query-lookout/domain"
	"github.com/KarnerTh/query-lookout/service/lookout"
)

func main() {
	l := lookout.New(domain.LookoutConfig{})
	l.Start()

	manageRunState()
}

// Keeps program running until SIGINT or SIGTERM
func manageRunState() {
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	fmt.Println("\n\nService stopped - see you soon 👋")
}
