package cron

import (
	"fmt"
)

func MainTask() {
	// 实现定时任务逻辑
	stopGitPull := gitPullTask()
	select {
	case <-stopGitPull:
		fmt.Println("停止 Git Pull")
		return
	default:
		fmt.Println("继续执行其他任务")
	}
	select {}
}
