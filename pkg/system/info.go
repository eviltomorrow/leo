package system

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/eviltomorrow/leo/pkg/tools"
)

var (
	Pid         = os.Getpid()
	Pwd         string
	LaunchTime  = time.Now()
	HostName    string
	OS          = runtime.GOOS
	Arch        = runtime.GOARCH
	RunningTime = func() string {
		return tools.FormatDuration(time.Since(LaunchTime))
	}
	IP string
)

func init() {
	path, err := os.Executable()
	if err != nil {
		panic(fmt.Errorf("get execute path failure, nest error: %v", err))
	}
	Pwd, err = filepath.Abs(path)
	if err != nil {
		panic(fmt.Errorf("get current folder failure, nest error: %v", err))
	}
	HostName, err = os.Hostname()
	if err != nil {
		panic(fmt.Errorf("get host name failure, nest error: %v", err))
	}
	IP, err = tools.GetLocalIP()
	if err != nil {
		panic(fmt.Errorf("get local ip failure, nest error: %v", err))
	}
}
