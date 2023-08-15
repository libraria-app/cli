/*
Copyright © 2023 libraria-app
*/
package main

import (
	"fmt"
	"os"

	"github.com/libraria-app/cli/commands"
	"github.com/libraria-app/cli/internal/librariacli"
	"github.com/libraria-app/cli/internal/utils/print"
)

func main() {
	if err := run(); err != nil {
		print.Error(fmt.Sprintf("error: %v \n", err))
		os.Exit(1)
	}
}

func run() error {
	lcli, err := librariacli.New()
	if err != nil {
		return err
	}

	return commands.NewRootCommand(lcli).Execute()
}
