package config

import (
	"context"
	"fmt"
	"reflect"
	"runtime"

	log "github.com/sirupsen/logrus"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/flags"
)

type ExporterConfig struct {
	Brand       string `config:"egym_brand"`
	Username    string `config:"egym_username"`
	Password    string `config:"egym_password"`
	MetricsPath string `config:"metricsPath"`
	BindAddress string `config:"bindAddress"`
	Port        uint16 `config:"port"`
}

func getDefaultConfig() *ExporterConfig {
	return &ExporterConfig{
		Brand:       "",
		Username:    "",
		Password:    "",
		BindAddress: "0.0.0.0",
		Port:        9391,
		MetricsPath: "/metrics",
	}
}

func Load() *ExporterConfig {
	loaders := []backend.Backend{
		env.NewBackend(),
		flags.NewBackend(),
	}

	loader := confita.NewLoader(loaders...)
	cfg := getDefaultConfig()
	err := loader.Load(context.Background(), cfg)
	if err != nil {
		panic(err)
	}

	cfg.show()

	return cfg
}

func (c ExporterConfig) show() {
	val := reflect.ValueOf(&c).Elem()
	log.Info("------------------------------------")
	log.Info("-      Exporter configuration      -")
	log.Info("------------------------------------")
	log.Info("Go version: ", runtime.Version())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		value := valueField.Interface()

		if typeField.Name == "Password" {
			value = "*************"
		}

		log.Info(fmt.Sprintf("%s : %v", typeField.Name, value))
	}
	log.Info("------------------------------------")
}
