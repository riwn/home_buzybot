package weather

import "github.com/riwn/home_discord_buzybot/pkg/domain/model/weather"

type Weather interface {
	FetchWeather() (weather.Weather, error)
}

type Discord interface {
	SendMessage(message string) error
}
