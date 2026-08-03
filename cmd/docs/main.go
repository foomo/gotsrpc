// Command docs generates the gotsrpc CLI reference as VitePress-friendly
// markdown from the cobra command tree.
//
// Usage:
//
//	go run ./cmd/docs --dir docs/reference/cli
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra/doc"

	"github.com/foomo/gotsrpc/v3/internal/cli"
)

func main() {
	dir := flag.String("dir", "docs/reference/cli", "output directory for the generated markdown")

	flag.Parse()

	if err := run(*dir); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}

func run(dir string) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("could not create output directory %s: %w", dir, err)
	}

	root := cli.NewRootCmd(cli.BuildInfo{Version: "dev", Commit: "n/a", BuildTime: "n/a"})

	// Only document the real commands: drop cobra's auto-generated completion
	// command so the reference stays focused on gotsrpc/generate/version.
	root.CompletionOptions.DisableDefaultCmd = true
	root.InitDefaultHelpCmd()

	if help, _, err := root.Find([]string{"help"}); err == nil && help != nil {
		root.RemoveCommand(help)
	}

	// filePrepender adds VitePress frontmatter so each page gets a title.
	filePrepender := func(filename string) string {
		name := filepath.Base(filename)
		name = strings.TrimSuffix(name, filepath.Ext(name))
		title := strings.ReplaceAll(name, "_", " ")

		return fmt.Sprintf("---\ntitle: %s\n---\n\n", title)
	}

	// linkHandler strips the .md suffix so cross-references resolve under
	// VitePress `cleanUrls`.
	linkHandler := func(name string) string {
		return "./" + strings.TrimSuffix(name, ".md")
	}

	if err := doc.GenMarkdownTreeCustom(root, dir, filePrepender, linkHandler); err != nil {
		return fmt.Errorf("could not generate markdown: %w", err)
	}

	return nil
}
