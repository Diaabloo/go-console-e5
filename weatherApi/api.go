package weatherApi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Cod     int    `json:"cod"`
	Message string `json:"message"`
	Name    string `json:"name"`
	Main    struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func GetWeather(city string) (*WeatherResponse, error) {
	// Charger .env
	_ = godotenv.Load()
	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=fr",
		city, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, err
	}

	return &weather, nil
}
