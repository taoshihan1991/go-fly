//+build !linux

package controller

import (
	"github.com/taoshihan1991/imaptool/config"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"time"
)

var osType = runtime.GOOS

const expireTime = 30 * 60

//检测权限文件是否过期,超过30分钟删除掉
func TimerSessFile() {
	go func() {
		for {
			time.Sleep(time.Second * 10)
			files, _ := filepath.Glob(config.Dir + "sess_*")
			for _, file := range files {
				fileInfo, _ := os.Stat(file)
				var createTime int64
				now := time.Now().Unix()
				if osType == "windows" {
					wFileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
					tNanSeconds := wFileSys.CreationTime.Nanoseconds() /// 返回的是纳秒
					createTime = tNanSeconds / 1e9                     ///秒
				}
				diffTime := now - createTime
				if diffTime > expireTime {
					os.Remove(file)
				}
			}
		}
	}()
}
