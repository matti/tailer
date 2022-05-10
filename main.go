package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
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

	for _, name := range os.Args[1:] {
		var fileName string
		var prefix string
		if strings.Contains(name, ":") {
			parts := strings.Split(name, ":")
			fileName = parts[0]
			prefix = parts[1]
		} else {
			fileName = name
			prefix = name
		}

		t, err := tail.TailFile(fileName, cfg)
		if err != nil {
			panic(err)
		}

		go func(prefix string, fileName string) {
			reader(prefix, fileName, t, cfg, output)
		}(prefix, fileName)
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

func reader(prefix string, fileName string, t *tail.Tail, cfg tail.Config, output chan string) {
	for line := range t.Lines {
		output <- (prefix + ": " + line.Text)
	}
}
