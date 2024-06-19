package cmd

import (
	"context"
	"fmt"
	"gin-web/config"
	"gin-web/internal/router"
	"gin-web/pkg/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var (
	port    string
	rootCmd = &cobra.Command{
		Use:          config.AppName,
		Short:        config.Description,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func completionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generate the autocompletion script for the specified shell",
	}
}

func init() {
	// 关闭官方completion命令
	completion := completionCommand()
	completion.Hidden = true
	rootCmd.AddCommand(completion)
	// 智能提示最小位数
	rootCmd.SuggestionsMinimumDistance = 1

	// 自定义log时间格式
	customTimeFormat := "2006-01-02 15:04:05.000"
	// 关闭默认的时间戳
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime)) // 2024-06-03 06:45:53.758 Server exiting
	// log.SetFlags(log.Llongfile) // 2024-06-03 06:44:22.090 C:/Users/OBY/Desktop/douyin-tool/cmd/root.go:102: Server exiting
	// 设置自定义的日志前缀
	log.SetPrefix(time.Now().Format(customTimeFormat) + " ")

	rootCmd.Flags().StringVarP(&port, "port", "p", "8080", "port")
}

func run() error {
	// 初始化路由
	r := router.InitRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// 处理非关闭服务器触发的错误
			log.Fatal("HTTP server ListenAndServe: ", err)
		}
	}()

	// 打印主机端口等信息
	fmt.Println("")
	fmt.Println("Server run at:")
	fmt.Printf("-    Local: http://localhost:%s \r\n", port)
	fmt.Printf("-  Network: http://%s:%s \r\n\n", util.GetLocaHonst(), port)
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", time.Now().Format("2006-01-02 15:04:05.000"))

	// 等待中断信号以优雅地关闭服务器
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞到接收到信号
	<-stop

	// 创建一个5秒的超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	return nil
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
