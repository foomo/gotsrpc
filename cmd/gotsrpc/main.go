package main

import (
	"fmt"
	"os"
)

func main() {
	if err := Execute(); err != nil {
		if logger != nil {
			logger.Error(err.Error())
		} else {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}

		os.Exit(1)
	}
}
