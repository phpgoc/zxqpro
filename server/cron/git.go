package cron

import (
	"fmt"
	"os/exec"

	"github.com/phpgoc/zxqpro/model/dao"

	"github.com/phpgoc/zxqpro/my_runtime"
	"github.com/phpgoc/zxqpro/utils"
	"github.com/robfig/cron/v3"
)

func gitPullTask() chan struct{} {
	// 查找project里是否有git地址，如果有就添加到定时任务里

	list := dao.ContainerInstance.ProjectDAO.GetAllGitPath()
	if len(list) > 0 {
		utils.LogWarn(fmt.Sprintf("git pull task list: %s", list))
	}
	for _, path := range list {
		if utils.IsGitRepository(path) {
			my_runtime.GitPathList.Add(path)
		}
	}

	// 实现 git pull 逻辑
	c := cron.New()
	_, err := c.AddFunc(fmt.Sprintf("*/%d * * * *", my_runtime.GitPullInterval), func() {
		//c := cron.New(cron.WithSeconds())
		//_, err := c.AddFunc("*/10 *  * * * *", func() {
		my_runtime.GitPathList.Each(func(value string) {
			cmd := exec.Command("git", "pull")
			cmd.Dir = value
			output, err := cmd.CombinedOutput()
			if err != nil {
				utils.LogError(fmt.Sprintf("git pull error: %s", err.Error()))
			} else {
				utils.LogInfo(fmt.Sprintf("git pull output: %s", string(output)))
			}
		})
	})
	if err != nil {
		utils.LogError(fmt.Sprintf("添加定时任务失败: %s", err.Error()))
		return nil
	}
	c.Start()

	// 创建控制通道
	stopChan := make(chan struct{})
	go func() {
		<-stopChan
		c.Stop()
		utils.LogInfo("定时任务已停止")
	}()
	return stopChan
}
