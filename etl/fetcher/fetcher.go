// Package allows to fetch all the APIs for data collection
package fetcher

import (
	"encoding/json"
	"etl/logger"
	"fmt"
	"io"
	"net/http"
)

var (
	Logger *logger.LogComponent
)

func init() {
	Logger = logger.Logger
}

// access site and return body
func get(url string) ([]byte, error) {
	// make http request
	resp, err := http.Get(url)
	if err != nil {
		Logger.LogError(err.Error())
		return nil, err
	}

	// close and read body
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		Logger.LogError(err.Error())
		return nil, err
	}

	return body, nil
}

func GetStations(lat string, long string, radius string, key string) (*IStation, error) {
	var stations IStation

	body, err := get(fmt.Sprintf("https://creativecommons.tankerkoenig.de/json/list.php?lat=%s&lng=%s&rad=%s&type=all&apikey=%s", lat, long, radius, key))
	if err != nil {
		Logger.LogError(err.Error())
		return &stations, err
	}

	// parse body response
	json.Unmarshal(body, &stations)
	return &stations, nil
}

// return array of parking garages
func GetParkingGarages() []IParkingGarage {
	body, err := get("https://api.parken-mannheim.de/")
	if err != nil {
		Logger.LogError(err.Error())
		return nil
	}

	// parse body response
	var parkingGarages []IParkingGarage
	json.Unmarshal(body, &parkingGarages)
	return parkingGarages
}

// return weather data of a city
func GetWeatherData(city string, key string) (*IWeather, error) {
	var weatherData IWeather

	body, err := get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?id=%s&units=metric&appid=%s", city, key))
	if err != nil {
		Logger.LogError(err.Error())
		return &weatherData, err
	}

	// parse body response
	json.Unmarshal(body, &weatherData)
	return &weatherData, nil
}

func GetTraffic(from string, to string, key string) (*ITraffic, error) {
	var trafficData ITraffic

	body, err := get(fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?units=metric&origins=%s&destinations=%s&departure_time=now&key=%s", from, to, key))
	if err != nil {
		Logger.LogError(err.Error())
		return &trafficData, err
	}

	// parse body response
	json.Unmarshal(body, &trafficData)
	return &trafficData, nil
}
