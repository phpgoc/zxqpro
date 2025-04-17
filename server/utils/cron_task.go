package utils

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/robfig/cron/v3"
)

func CronTask() {
	// todo  查找project里是否有git地址，如果有就添加到定时任务里

	// 实现定时任务逻辑
	stopGitPull := gitPull()
	select {
	case <-stopGitPull:
		fmt.Println("停止 Git Pull")
		return
	default:
		fmt.Println("继续执行其他任务")
	}
	select {}
}

type stringSet struct {
	set sync.Map
}

func (s *stringSet) Add(value string) {
	s.set.Store(value, struct{}{})
}

func (s *stringSet) Remove(value string) {
	s.set.Delete(value)
}

func (s *stringSet) Each(f func(value string)) {
	s.set.Range(func(key, value interface{}) bool {
		f(key.(string))
		return true
	})
}

var GitPathList stringSet

func gitPull() chan struct{} {
	// 实现 git pull 逻辑
	c := cron.New()
	_, err := c.AddFunc("*/30 * * * *", func() {
		//c := cron.New(cron.WithSeconds())
		//_, err := c.AddFunc("*/10 *  * * * *", func() {
		GitPathList.Each(func(value string) {
			cmd := exec.Command("git", "pull")
			cmd.Dir = value
			output, err := cmd.CombinedOutput()
			if err != nil {
				LogError(fmt.Sprintf("git pull error: %s", err.Error()))
			} else {
				LogInfo(fmt.Sprintf("git pull output: %s", string(output)))
			}
		})
	})
	if err != nil {
		LogError(fmt.Sprintf("添加定时任务失败: %s", err.Error()))
		return nil
	}
	c.Start()

	// 创建控制通道
	stopChan := make(chan struct{})
	go func() {
		<-stopChan
		c.Stop()
		LogInfo("定时任务已停止")
	}()
	return stopChan
}
