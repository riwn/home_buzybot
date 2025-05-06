package main

import (
	"fmt"

	"github.com/riwn/home_buzybot/pkg/registry"
	"github.com/riwn/home_buzybot/pkg/setting"
	"github.com/riwn/home_buzybot/pkg/uc/weather"
)

func main() {
	registry := registry.New(setting.Get())
	uc := weather.NewWeatherNotification(registry)
	if err := uc.SendWeatherNotification(); err != nil {
		fmt.Println(err)
		return
	}
}
