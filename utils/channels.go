package utils

import (
	"os"
	"os/signal"
)

type MapperFunc[T any, U any] func(T) U

// WrapChannel wraps an input channel into another channel using a mapper function
func WrapChannel[T any, U any](input <-chan T, mapper MapperFunc[T, U]) <-chan U {
	output := make(chan U)

	go func() {
		for item := range input {
			output <- mapper(item)
		}
		close(output)
	}()

	return output
}

func SetUpProcessInterrupt(onExitFunc func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		if onExitFunc != nil {
			onExitFunc()
		}
		os.Exit(1)
	}()
}
