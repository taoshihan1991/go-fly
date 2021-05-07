package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-fly",
	Short: "go-fly",
	Long:  `简洁快速的GO语言WEB在线客服 https://gofly.sopans.com`,
	Args:  args,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func args(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {

		return errors.New("至少需要一个参数!")
	}
	return nil
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(stopCmd)
}
