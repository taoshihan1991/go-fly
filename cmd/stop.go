package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strings"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "停止http服务",
	Run: func(cmd *cobra.Command, args []string) {
		pids, err := ioutil.ReadFile("gofly.sock")
		if err != nil {
			return
		}
		pidSlice := strings.Split(string(pids), ",")
		var command *exec.Cmd
		for _, pid := range pidSlice {
			if runtime.GOOS == "windows" {
				command = exec.Command("taskkill.exe", "/f", "/pid", pid)
			} else {
				command = exec.Command("kill", pid)
			}
			command.Start()
		}
	},
}
