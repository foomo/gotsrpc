package main

import (
	"fmt"
	"os"

	"github.com/foomo/gotsrpc/v3/internal/cli"
)

// Build-time metadata, injected via -ldflags (see Makefile).
var (
	version        = "v3-dev"
	commitHash     = "n/a"
	buildTimestamp = "n/a"
)

func main() {
	err := cli.Execute(cli.BuildInfo{
		Version:   version,
		Commit:    commitHash,
		BuildTime: buildTimestamp,
	})
	if err != nil {
		if logger := cli.Logger(); logger != nil {
			logger.Error(err.Error())
		} else {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}

		os.Exit(1)
	}
}
