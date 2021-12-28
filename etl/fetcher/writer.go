package fetcher

import (
	"etl/writer"
	"fmt"
	"path/filepath"
	"strconv"
)

func (p IParkingGarage) Write(filename string, datetime string) {
	var err = writer.WriteMetaData(filepath.Join("parking", "meta"), p.Identifier, []string{"title", "provider", "slots"}, []string{p.Title, p.Provider, p.Slots})
	if err != nil {
		Logger.LogError(fmt.Sprintf("Error Writing Meta Data for %s: %s", p.Identifier, err.Error()))
	}

	err = writer.WriteValueData("parking", fmt.Sprintf("%s.csv", filename), []string{"date", "identifier", "free_slots"}, [][]string{{datetime, p.Identifier, p.FreeSlots}})
	if err != nil {
		Logger.LogError(fmt.Sprintf("Error Writing Value Data for %s: %s", p.Identifier, err.Error()))
	}
}

func (s Stations) Write(filename string, datetime string) {
	var err = writer.WriteMetaData(filepath.Join("stations", "meta"), s.ID, []string{"name", "brand", "street", "place"}, []string{s.Name, s.Brand, s.Street, s.Place})
	if err != nil {
		Logger.LogError(fmt.Sprintf("Error Writing Meta Data for %s: %s", s.ID, err.Error()))
	}

	err = writer.WriteValueData("stations", fmt.Sprintf("%s.csv", filename),
		[]string{"date", "id", "diesel", "e5", "e10", "isOpen"},
		[][]string{{datetime, s.ID, strconv.FormatFloat(s.Diesel, 'f', -1, 64), strconv.FormatFloat(s.E5, 'f', -1, 64), strconv.FormatFloat(s.E10, 'f', -1, 64), strconv.FormatBool(s.IsOpen)}})
	if err != nil {
		Logger.LogError(fmt.Sprintf("Error Writing Value Data for %s: %s", s.ID, err.Error()))
	}
}

func (w IWeather) Write(filename string, datetime string) {
	var err = writer.WriteMetaData(filepath.Join("weather", "meta"), strconv.Itoa(w.ID),
		[]string{"id", "lon", "lat", "base", "timezone", "name"},
		[]string{strconv.Itoa(w.ID), strconv.FormatFloat(w.Coord.Lon, 'f', -1, 64), strconv.FormatFloat(w.Coord.Lat, 'f', -1, 64), w.Base, strconv.Itoa(w.Timezone), w.Name})

	if err != nil {
		Logger.LogError(fmt.Sprintf("Error Writing Meta Data for %s: %s", w.ID, err.Error()))
	}

	err = writer.WriteValueData("weather", fmt.Sprintf("%s.csv", filename),
		[]string{"date", "id", "temp", "feels_like", "temp_min", "temp_max", "pressure", "humidity", "sea_level", "grnd_level", "visibility", "wind_speed", "wind_deg"},
		[][]string{{datetime, strconv.Itoa(w.ID), strconv.FormatFloat(w.Main.Temp, 'f', -1, 64), strconv.FormatFloat(w.Main.FeelsLike, 'f', -1, 64),
			strconv.FormatFloat(w.Main.TempMin, 'f', -1, 64),
			strconv.FormatFloat(w.Main.TempMax, 'f', -1, 64), strconv.Itoa(w.Main.Pressure), strconv.Itoa(w.Main.Humidity),
			strconv.Itoa(w.Main.SeaLevel), strconv.Itoa(w.Main.GrndLevel), strconv.Itoa(w.Visibility), strconv.FormatFloat(w.Wind.Speed, 'f', -1, 64),
			strconv.Itoa(w.Wind.Deg)}})
	if err != nil {
		Logger.LogError(fmt.Sprintf("Error Writing Value Data for %s: %s", w.ID, err.Error()))
	}
}

func (t ITraffic) Write(filename string, datetime string, lat string, lng string) {
	var err = writer.WriteMetaData(filepath.Join("traffic", "meta"), fmt.Sprintf("%s-%s", lat, lng), []string{"lat", "lng"}, []string{lat, lng})
	if err != nil {
		Logger.LogError(fmt.Sprintf("Error Writing Meta Data for %s: %s", fmt.Sprintf("%s-%s", lat, lng), err.Error()))
	}

	var distance, duration, durationInTraffic int = 0, 0, 0
	for _, row := range t.Rows {
		for _, element := range row.Elements {
			distance += element.Distance.Value
			duration += element.Duration.Value
			durationInTraffic += element.DurationInTraffic.Value
		}
	}

	err = writer.WriteValueData("traffic", fmt.Sprintf("%s.csv", filename),
		[]string{"date", "id", "distance", "duration", "durationInTraffic"},
		[][]string{{datetime, fmt.Sprintf("%s-%s", lat, lng), strconv.Itoa(distance), strconv.Itoa(duration), strconv.Itoa(durationInTraffic)}})
	if err != nil {
		Logger.LogError(fmt.Sprintf("Error Writing Value Data for %s: %s", fmt.Sprintf("%s-%s", lat, lng), err.Error()))
	}
}
