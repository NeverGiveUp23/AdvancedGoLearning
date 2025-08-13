package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

type Container[T any] struct{ value T }

func main() {
	fmt.Println(KillServer("server.pid"))
}

func KillServer(pidFile string) error {
	if _, err := os.Create(pidFile); err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	file, err := os.Open(pidFile)
	if err != nil {
		log.Printf("error opening file: %v", err)
		return err
	}

	err = os.WriteFile(pidFile, []byte("Original Message"), os.ModeAppend)
	if err != nil {
		return fmt.Errorf("failed tp write to file: %w", err)
	}

	// Error check option 1 by anonymous defer function
	// Defer works at a function level
	// Defer are executed in reverse order (stack, LIFO)
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("error closing file: %v", err)
		}
	}()

	var pid Container[string]
	if _, err := fmt.Fscanf(file, "%s", &pid.value); err != nil {
		return fmt.Errorf("%q - bad pid: %s", pidFile, err)
	}

	slog.Info("killing", "pid", pid)

	return nil
}
