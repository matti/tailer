package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hpcloud/tail"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	cfg := tail.Config{
		Follow:    true,
		ReOpen:    true,
		MustExist: false,
		Logger:    tail.DiscardingLogger,
	}

	output := make(chan string)

	for _, fileName := range os.Args[1:] {
		t, err := tail.TailFile(fileName, cfg)
		if err != nil {
			panic(err)
		}

		go func(fileName string) {
			reader(fileName, t, cfg, output)
		}(fileName)
	}

	for {
		select {
		case <-sigs:
			os.Exit(0)
		case line := <-output:
			fmt.Println(line)
		}
	}
}

func reader(fileName string, t *tail.Tail, cfg tail.Config, output chan string) {
	for line := range t.Lines {
		output <- (fileName + ": " + line.Text)
	}
}
