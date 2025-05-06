package weather

import "fmt"

type (
	Weather struct {
		PlaceName       PlaceName
		WeatherForecast WeatherForecast
		DayTemps        DayTemps
		Humidity        Humidity
	}

	Temperature float64
	Humidity    int
	PlaceName   string
	DayTemps    struct {
		Now       Temperature
		FeelsLike Temperature
		Min       Temperature
		Max       Temperature
	}
	WeatherForecast struct {
		MainWeather string
		Description string
	}
)

func ReconstructWeather(
	placeName PlaceName,
	weatherForecast WeatherForecast,
	dayTemps DayTemps,
	humidity Humidity,
) Weather {
	return Weather{
		PlaceName:       placeName,
		WeatherForecast: weatherForecast,
		DayTemps:        dayTemps,
		Humidity:        humidity,
	}
}

func (w Weather) BuildDiscordMessage() string {
	return fmt.Sprintf(
		"☀️ **天気情報**\n"+
			"🌡️ 気温: %.1f℃（体感気温: %.1f℃）\n"+
			"📈 最高: %.1f℃　📉 最低: %.1f℃\n"+
			"🌥️ 天気: %s（%s）\n"+
			"💧 湿度: %d%%",
		w.DayTemps.Now,
		w.DayTemps.FeelsLike,
		w.DayTemps.Max,
		w.DayTemps.Min,
		w.WeatherForecast.MainWeather,
		w.WeatherForecast.Description,
		w.Humidity,
	)
}
