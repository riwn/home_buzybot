package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/riwn/home_buzybot/pkg/domain/model/weather"
)

type Weather struct {
	url    string
	lat    string
	lon    string
	apiKey string
}

func NewWeather(url, lat, lon, apiKey string) Weather {
	return Weather{
		url:    url,
		lat:    lat,
		lon:    lon,
		apiKey: apiKey,
	}
}

func (w Weather) FetchWeather() (weather.Weather, error) {
	url := fmt.Sprintf(
		"%s?lat=%s&lon=%s&appid=%s&units=metric&lang=ja",
		w.url,
		w.lat,
		w.lon,
		w.apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return weather.Weather{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return weather.Weather{}, err
	}
	if resp.StatusCode != 200 {
		return weather.Weather{}, fmt.Errorf("weather API error: %d, %s", resp.StatusCode, body)
	}
	var data WeatherResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return weather.Weather{}, err
	}

	return data.ToModel(), nil
}
