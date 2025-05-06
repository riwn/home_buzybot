package setting

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/Masterminds/sprig/v3"
	"gopkg.in/yaml.v3"
)

type Config struct {
	OpenWeather struct {
		URL       string `yaml:"url"`
		APIKey    string `yaml:"api_key"`
		Latitude  string `yaml:"latitude"`
		Longitude string `yaml:"longitude"`
	} `yaml:"open_weather"`
	Discord struct {
		WebhookURL string `yaml:"webhook_url"`
	} `yaml:"discord"`
}

func (c Config) String() string {
	return "OpenWeather URL: " + c.OpenWeather.URL + "\n" +
		"OpenWeather API Key: " + c.OpenWeather.APIKey + "\n" +
		"OpenWeather Latitude: " + c.OpenWeather.Latitude + "\n" +
		"OpenWeather Longitude: " + c.OpenWeather.Longitude + "\n" +
		"Discord Webhook URL: " + c.Discord.WebhookURL
}

//go:embed config.yaml
var file embed.FS

var globalConfig Config

func init() {
	config, err := LoadConfig("config.yaml")
	if err != nil {
		panic("設定ファイルの読み込みに失敗しました: " + err.Error())
	}
	globalConfig = config

	fmt.Println(config)
}

func Get() Config {
	return globalConfig
}

func LoadConfig(path string) (Config, error) {
	tmpl, err := file.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	t, err := template.New(filepath.Base(path)).
		Funcs(sprig.TxtFuncMap()).
		Parse(string(tmpl))

	if err != nil {
		return Config{}, err
	}

	var b bytes.Buffer
	if err := t.Execute(&b, nil); err != nil {
		return Config{}, err
	}

	var config Config
	if err := yaml.Unmarshal(b.Bytes(), &config); err != nil {
		panic(err)
	}
	return config, nil
}
