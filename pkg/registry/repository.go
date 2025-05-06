package registry

import (
	"github.com/riwn/home_discord_buzybot/pkg/infra/discord"
	"github.com/riwn/home_discord_buzybot/pkg/infra/weather"
	"github.com/riwn/home_discord_buzybot/pkg/setting"
)

type Repository struct {
	config setting.Config
}

func NewRepository(config setting.Config) Repository {
	return Repository{
		config: config,
	}
}

func (r Repository) NewWeatherRepository() weather.Weather {
	return weather.NewWeather(
		r.config.OpenWeather.URL,
		r.config.OpenWeather.Latitude,
		r.config.OpenWeather.Longitude,
		r.config.OpenWeather.APIKey,
	)
}

func (r Repository) NewDiscordRepository() discord.Discord {
	return discord.NewDiscord(r.config.Discord.WebhookURL)
}
