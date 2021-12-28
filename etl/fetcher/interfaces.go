package fetcher

type IStation struct {
	Ok       bool       `json:"ok"`
	Stations []Stations `json:"stations"`
}

type Stations struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Brand  string  `json:"brand"`
	Street string  `json:"street"`
	Place  string  `json:"place"`
	Diesel float64 `json:"diesel"`
	E5     float64 `json:"e5"`
	E10    float64 `json:"e10"`
	IsOpen bool    `json:"isOpen"`
}

type IParkingGarage struct {
	Identifier string `json:"identifier"`
	WpPostID   string `json:"wp_post_id,omitempty"`
	Title      string `json:"title"`
	Provider   string `json:"provider"`
	FreeSlots  string `json:"free_slots"`
	Updated    string `json:"updated"`
	Sort       string `json:"sort"`
	Slots      string `json:"slots,omitempty"`
}

type ITraffic struct {
	DestinationAddresses []string `json:"destination_addresses"`
	OriginAddresses      []string `json:"origin_addresses"`
	Rows                 []struct {
		Elements []struct {
			Distance struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"distance"`
			Duration struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"duration"`
			DurationInTraffic struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"duration_in_traffic"`
			Status string `json:"status"`
		} `json:"elements"`
	} `json:"rows"`
	Status string `json:"status"`
}

type IWeather struct {
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

	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`

	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`

	Dt int `json:"dt"`

	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`

	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}
