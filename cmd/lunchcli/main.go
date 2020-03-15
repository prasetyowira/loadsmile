package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/prasetyowira/loadsmile/internal/app/lunchcli"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "lunchcli",
		Short: "Lunch CLI manages Recipes and Ingredients.",
	}

	lunchcli.Configure(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
