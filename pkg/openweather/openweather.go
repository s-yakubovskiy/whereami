package openweather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherData struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int64 `json:"dt"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type OpenWeatherReport struct {
	// http://api.openweathermap.org/data/2.5/weather
	url     string
	api_key string
}

func NewOpenWeatherReport(url, apikey string) (*OpenWeatherReport, error) {
	return &OpenWeatherReport{
		url:     url,
		api_key: apikey,
	}, nil
}

func (s *OpenWeatherReport) FetchWeather(lat, lon float64) (*WeatherData, error) {
	url := fmt.Sprintf("%s?lat=%f&lon=%f&units=metric&appid=%s", s.url, lat, lon, s.api_key)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherData *WeatherData
	err = json.Unmarshal(jsonData, &weatherData)
	if err != nil {
		return nil, err
	}

	return weatherData, nil
}
