package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"syscall"

	filemonitor "github.com/Fromsko/gouitls/monitor"
)

func HotReload() {
	// 用于存储上次执行的进程
	var lastCmd *exec.Cmd

	// 创建文件监控组件
	monitor := filemonitor.NewFileMonitor("./dist", func(fm *filemonitor.FileMonitor) {
		log.Println("文件重载...")

		// 尝试终止上次执行的进程
		if lastCmd != nil && lastCmd.Process != nil {
			var err error
			if runtime.GOOS == "windows" {
				// Windows
				err = lastCmd.Process.Kill()
			} else {
				// Unix-like systems
				err = lastCmd.Process.Signal(syscall.SIGTERM)
			}
			if err != nil {
				log.Println("Error terminating previous process:", err)
			}
		}

		// 创建新进程
		cmd_ := exec.Command(
			"statik", "-src=dist",
		)
		_ = cmd_.Start()

		// 编译
		cmd := exec.Command(
			"make clean && make",
		)
		// 设置新进程的工作目录为当前目录
		_ = cmd.Start()
		cmd.Dir, _ = os.Getwd()

		// 在单独的 goroutine 中等待新进程完成
		go func() {
			err := cmd.Wait()
			if err != nil {
				log.Println("Error waiting for restart:", err)
			}
		}()

		// 更新 lastCmd
		lastCmd = cmd
	})

	// 启动文件监控服务
	go monitor.Start()

	// 你的主程序逻辑
	fmt.Println("程序正在运行！")

	// 让程序一直运行
	select {}
}

func main() {
	HotReload()
}
