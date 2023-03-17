package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/KarnerTh/query-lookout/orchestration"
)

func main() {
	orchestration.Setup()

	// Keep program running until SIGINT or SIGTERM
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	fmt.Println("Service stopped - see you soon 👋")
}
