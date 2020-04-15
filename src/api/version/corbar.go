package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

/**
 * description: 返回一个用于打印version信息的命令
 */
func Command() *cobra.Command {
	var short bool
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prints out build version information",
		Run: func(cmd *cobra.Command, args []string) {
			if short {
				fmt.Println(Info)
			} else {
				fmt.Println(Info.LongForm())
			}
		},
	}
	cmd.PersistentFlags().BoolVarP(&short, "short", "s", short, "Displays a short form of the version information")
	return cmd
}
