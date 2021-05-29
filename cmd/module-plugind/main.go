package main

import (
	"os"

	"github.com/mconcat/module-plugin/cmd/module-plugind/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
