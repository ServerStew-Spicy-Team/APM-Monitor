package apm_monitor

import (
	"APM-Monitor/internal/pkg/log"
	"APM-Monitor/pkg/kafka"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		//home, err := os.UserHomeDir()
		//cobra.CheckErr(err)

		// Search config in home directory with name ".APM-Monitor" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("apm-monitor.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	viper.SetEnvPrefix("APM-MONITOR")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	//if err := godotenv.Load(".env"); err != nil {
	//	log.Fatalw("Error loading .env file")
	//}
}

func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}

func initProducer() error {
	kafkaOpts := &kafka.KafkaOptions{
		ProducerReturnSuccess: viper.GetBool("kafka.producer-return-success"),
		ProducerReturnErr:     viper.GetBool("kafka.producer-return-err"),
		Brokers:               viper.GetStringSlice("kafka.brokers"),
	}
	ins, err := kafka.NewProducer(kafkaOpts)
	if err != nil {
		return err
	}

	kafka.StoreProducer(ins)
	log.Infow("kafka producer running")
	return nil
}
