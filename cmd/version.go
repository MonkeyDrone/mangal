package cmd

import (
	"fmt"
	"github.com/metafates/mangal/common"
	"github.com/metafates/mangal/style"
	"github.com/spf13/cobra"
	"strings"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  fmt.Sprintf("Shows %s version", common.Mangal),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %s\n", strings.ToLower(common.Mangal), style.Accent.Render(common.Version))
	},
}

func init() {
	mangalCmd.AddCommand(versionCmd)
}