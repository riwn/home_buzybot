package weather

import "github.com/riwn/home_discord_buzybot/pkg/registry"

type WeatherNotification struct {
	weather Weather
	discord Discord
}

func NewWeatherNotification(rgst registry.Registry) WeatherNotification {
	return WeatherNotification{
		weather: rgst.Repository().NewWeatherRepository(),
		discord: rgst.Repository().NewDiscordRepository(),
	}
}

func (w WeatherNotification) SendWeatherNotification() error {
	weatherData, err := w.weather.FetchWeather()
	if err != nil {
		return err
	}

	if err := w.discord.SendMessage(weatherData.BuildDiscordMessage()); err != nil {
		return err
	}

	return nil
}
