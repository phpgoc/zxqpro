package utils

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

var (
	logFilePath = getEnv("LOG_FILE_PATH", "logs/zxqpro.log")
	useLogFile  = getEnv("USE_LOG_FILE", "0")
	CookieName  = "zxqpro_cookie"
	version     = "v0.1.0"
)

var (
	RedisAddr             = ""
	GitPullInterval uint8 = 10 // Minutes
)

var (
	versionCommand = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of zxqpro",
		Long:  `All software has versions. This is zxqpro's`,
		Run: func(cmd *cobra.Command, args []string) {
			// Print the version number
			cmd.Println("zxqpro version:", version)
			os.Exit(0)
		},
	}
	rootCmd = &cobra.Command{
		Use:   "zxqpro",
		Short: "zxqpro application",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// 在这里添加你想要在执行任何子命令前执行的代码
			fmt.Println("Pre-run logic...")
		},
		Run: func(cmd *cobra.Command, args []string) {
			// 执行根命令的逻辑，这里可以添加你需要的操作
			if GitPullInterval < 1 || GitPullInterval > 59 {
				cmd.Println("Git pull interval must be between 1 and 60")
				os.Exit(1)
			}
		},
	}
)

func InitCobra() {
	rootCmd.AddCommand(versionCommand)
	rootCmd.Flags().Uint8VarP(&GitPullInterval, "git_pull_interval", "g", 10, "Git pull interval in minutes 1-59")

	rootCmd.Flags().StringVarP(&RedisAddr, "redis_addr", "r", "", "Redis address,default use go-cache if empty")
	err := rootCmd.Execute()
	for _, arg := range os.Args[1:] {
		if arg == "help" || arg == "--help" || arg == "-h" {
			os.Exit(0)
		}
	}
	if err != nil {
		// 解析参数出错时显示帮助信息并退出
		os.Exit(1)
	}
}
