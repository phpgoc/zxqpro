package utils

import (
	"errors"
	"os"

	"gorm.io/gorm/logger"

	"github.com/spf13/cobra"
)

var (
	GinLogWriter    = os.Stdout
	GormLogWriter   = os.Stdout
	SelfLogWriter   = os.Stdout
	CookieName      = "zxqpro_cookie"
	version         = "v0.1.0"
	gormLogLevelMap = map[string]logger.LogLevel{
		"s": logger.Silent,
		"i": logger.Info,
		"w": logger.Warn,
		"e": logger.Error,
	}
	GormLogLevel logger.LogLevel = gormLogLevelMap["w"]
)

var (
	RedisAddr                = ""
	GitPullInterval    uint8 = 10 // Minutes
	Port                     = 8080
	GormLogLevelString       = "d"
	GinDebugModel            = true
	ginLogFile               = ""
	gormLogFile              = ""
	selfLogFile              = ""
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
			// fmt.Println("Pre-run logic...")
		},
		Run: func(cmd *cobra.Command, args []string) {
			// 执行根命令的逻辑，这里可以添加你需要的操作
			if GitPullInterval < 1 || GitPullInterval > 59 {
				cmd.Println("Git pull interval must be between 1 and 60")
				os.Exit(1)
			}
			if _, ok := gormLogLevelMap[GormLogLevelString]; !ok {
				cmd.Println("Gorm log level must be one of: i, w, e, s")
				os.Exit(1)
			} else {
				GormLogLevel = gormLogLevelMap[GormLogLevelString]
			}
			if gormLogFile != "" {
				if err := isValidPathAndSetWriter(gormLogFile, &GormLogWriter); err != nil {
					cmd.Println("gorm_log error path")
					os.Exit(1)
				}
			}
			if ginLogFile != "" {
				if err := isValidPathAndSetWriter(ginLogFile, &GinLogWriter); err != nil {
					cmd.Println("gin_log error path")
					os.Exit(1)
				}
			}
			if selfLogFile != "" {
				if err := isValidPathAndSetWriter(selfLogFile, &SelfLogWriter); err != nil {
					cmd.Println("self_log error path")
					os.Exit(1)
				}
			}
		},
	}
)

func isValidPathAndSetWriter(path string, writer **os.File) error {
	if path == "" {
		return errors.New("empty path")
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	*writer = file

	return nil
}

func InitCobra() {
	rootCmd.AddCommand(versionCommand)
	rootCmd.Flags().Uint8VarP(&GitPullInterval, "git_pull_interval", "g", 10, "Git pull interval in minutes 1-59")

	rootCmd.Flags().StringVarP(&RedisAddr, "redis_addr", "r", "", "Redis address,default use go-cache if empty")
	rootCmd.Flags().IntVarP(&Port, "port", "p", 8080, "Port to listen on")
	rootCmd.Flags().StringVarP(&GormLogLevelString, "gorm_log_level", "l", "w", "Gorm log level: i, w, e, s")
	rootCmd.Flags().BoolVarP(&GinDebugModel, "debug", "d", false, "Gin debug model, default is ReleaseMode")
	rootCmd.Flags().StringVar(&ginLogFile, "gin_log", "", "Gin log file")
	rootCmd.Flags().StringVar(&gormLogFile, "gorm_log", "", "Gorm log file")
	rootCmd.Flags().StringVar(&selfLogFile, "self_log", "", "Self define log file")
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
