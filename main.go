package main

import (
	"flag"
	"fmt"
	"github.com/open-falcon/recivers/g"
	"github.com/open-falcon/recivers/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	flag.Parse()


	g.ParseConfig(*cfg)
	go http.Start()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println()
		os.Exit(0)
	}()

	select {}
}
