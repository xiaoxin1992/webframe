package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"path"
	"syscall"
	"webframe/conf"
	"webframe/pkg/logger"
	"webframe/version"
	"webframe/pkg/app"
)

var (
	ver        bool
	configPath string
	rootPath   string
)

var rootCmd = &cobra.Command{
	Use:   "projectName",
	Short: "projectName说明",
	Long:  `projectName描述`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if ver {
			fmt.Println(version.FullVersion())
			return nil
		}
		rootPath = path.Dir(os.Args[0])
		if !path.IsAbs(configPath) {
			configPath = path.Join(rootPath, configPath)
		}
		// 初始化配置文件
		if err := conf.NewConfig(configPath); err != nil {
			return err
		}
		// 初始化日志配置
		if err := logger.NewLogger(rootPath); err != nil {
			return err
		}
		// 初始化全局APP的config
		if err := app.InitAllApp(); err != nil {
			return err
		}
		// 启动服务
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
		svr := newService()
		go svr.waitSign(ch)
		if err := svr.Start(); err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "etc/config.toml", "配置文件路径")
	rootCmd.PersistentFlags().BoolVarP(&ver, "version", "v", false, "the webframe version")
}
