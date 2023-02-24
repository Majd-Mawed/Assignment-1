package handler

type Coordinates struct {
	Latitude  float64 `json:"lat,omitempty"`
	Longitude float64 `json:"lon,omitempty"`
}

type Location struct {
	Name        string      `json:"name"`
	Postcode    string      `json:"code"`
	Country     string      `json:"country,omitempty"` // 'omitempty' suppresses field in JSON output if it is empty
	Geolocation Coordinates `json:"location"`
}

type UNI struct {
	Name     string   `json:"name"`
	Country  string   `json:"country"`
	Webpages []string `json:"web_pages"`
	Isocode  string   `json:"alpha_two_code"`
}
