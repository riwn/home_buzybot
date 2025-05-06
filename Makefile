build-raspberry: ## Build for Raspberry Pi (ARMv7)
	GOOS=linux GOARCH=arm GOARM=7 go build -o weather-bot ./cmd/send_weather/main.go

.PHONY: build