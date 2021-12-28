package translator

import (
	"encoding/json"
	"etl/fetcher"
	"etl/logger"
	"os"
	"time"
)

var (
	Logger   *logger.LogComponent
	filename string
	datetime string
)

func init() {
	Logger = logger.Logger
	var currTime = time.Now()
	datetime = currTime.Format(time.RFC3339)
	filename = currTime.Format("2006-01-02")
}

func FetchParking() {
	Logger.LogInfo("Started Fetching for Parking")
	for _, parking := range fetcher.GetParkingGarages() {
		parking.Write(filename, datetime)
	}
	Logger.LogInfo("Finished Fetching for Parking")
}

func FetchStation() {
	Logger.LogInfo("Started Fetching for Petrol Stations")
	var stationAPI string = os.Getenv("STATIONAPI")
	if stationAPI == "" {
		Logger.LogError("STATIONAPI NOT DEFINED")
		os.Exit(1)
	}
	// parse list of locations
	var listOfLocations [][]string
	json.Unmarshal([]byte(os.Getenv("STATIONLOCATIONS")), &listOfLocations)
	if len(listOfLocations) == 0 {
		Logger.LogError("List of Locations for Stations is Empty")
	}
	// fetch stations for all locations
	for _, element := range listOfLocations {
		fetchedData, err := fetcher.GetStations(element[0], element[1], element[2], stationAPI)
		if err == nil {
			for _, fetchedStation := range fetchedData.Stations {
				fetchedStation.Write(filename, datetime)
			}
		}
	}
	Logger.LogInfo("Finished Fetching for Petrol Stations")
}

func FetchTraffic() {
	Logger.LogInfo("Started Fetching for Traffic")
	var trafficAPI string = os.Getenv("TRAFFICAPI")
	if trafficAPI == "" {
		Logger.LogError("TRAFFICAPI NOT DEFINED")
		os.Exit(1)
	}
	// parse list of directions
	var listOfDirections [][]string
	json.Unmarshal([]byte(os.Getenv("TRAFFICLIST")), &listOfDirections)
	if len(listOfDirections) == 0 {
		Logger.LogError("List of Directions for Traffic is Empty")
	}
	// fetch traffic for all directions
	for _, element := range listOfDirections {
		fetchedData, err := fetcher.GetTraffic(element[0], element[1], trafficAPI)
		if err == nil {
			fetchedData.Write(filename, datetime, element[0], element[1])
		}
	}
	Logger.LogInfo("Finished Fetching for Traffic")
}

func FetchWeather() {
	Logger.LogInfo("Started Fetching for Weather")
	// get weather api key and exit if it wasn't defined
	var weatherAPI string = os.Getenv("WEATHERAPI")
	if weatherAPI == "" {
		Logger.LogError("WEATHERAPI NOT DEFINED")
		os.Exit(1)
	}

	// parse list of cities
	var listOfCities []string
	json.Unmarshal([]byte(os.Getenv("WEATHERCITYLIST")), &listOfCities)
	if len(listOfCities) == 0 {
		Logger.LogError("List of Cities for Weather is Empty")
	}

	// fetch weather for all cities
	for _, element := range listOfCities {
		fetchedData, err := fetcher.GetWeatherData(element, weatherAPI)
		if err == nil {
			fetchedData.Write(filename, datetime)
		}
	}
	Logger.LogInfo("Finished Fetching for Weather")
}
