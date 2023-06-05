/*
Copyright © 2023 NAME HERE <784312513@qq.com>
*/
package apm_monitor

import (
	"APM-Monitor/internal/apm-monitor/scheduler"
	"APM-Monitor/internal/pkg/known"
	"APM-Monitor/internal/pkg/log"
	"APM-Monitor/pkg/kafka"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var cfgFile string

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "APM-Monitor",
	Short: "A monitor for each server",

	RunE: func(cmd *cobra.Command, args []string) error {
		log.Init(logOptions())

		defer log.Sync() // Sync 将缓存中的日志刷新到磁盘文件中
		return run()
	},
	Args: func(cmd *cobra.Command, args []string) error {
		for _, arg := range args {
			if len(arg) > 0 {
				return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
			}
		}

		return nil
	},
}

func run() error {

	log.Infow("logger is running", "level:", viper.GetString("log.level"))

	err := initProducer()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(4)
	ctx, cancel := context.WithCancel(context.Background())
	//业务代码
	go scheduler.Schedule(ctx, known.CPU, &wg)
	go scheduler.Schedule(ctx, known.MEMORY, &wg)
	go scheduler.Schedule(ctx, known.DISK, &wg)
	go scheduler.Schedule(ctx, known.NETWORK, &wg)

	//graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infow("Server exiting")
	cancel()
	wg.Wait()
	func() {
		err := kafka.Pro().P.Close()
		if err != nil {
			log.Fatalw("kafka shutdown error,", "err", err)
		}
		log.Infow("Producer successfully shutdown")
	}()

	return nil
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the apm-monitor configuration file. Empty string for no configuration file")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
