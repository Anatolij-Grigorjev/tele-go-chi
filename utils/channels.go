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

type ExitFunc func()

var exitFuncs []ExitFunc = make([]ExitFunc, 0)

func AddOnExitFunc(onExit ExitFunc) {
	exitFuncs = append(exitFuncs, onExit)
}

// Sets up receiver for OS interrupt signal.
//
// To run code when that happens, add functions via the AddOnExitFunc(func()) process
func SetUpProcessInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		for _, actions := range exitFuncs {
			actions()
		}
		os.Exit(1)
	}()
}
