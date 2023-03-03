package handler

type UNI struct {
	Name     string   `json:"name"`
	Country  string   `json:"country"`
	Webpages []string `json:"web_pages"`
	Isocode  string   `json:"alpha_two_code"`
}

type NABUNI struct {
	Name struct {
		Name string `json:"common"`
	} `json:"name"`
	Languages map[string]interface{} `json:"languages"`
	Map       map[string]interface{} `json:"maps"`
}

type NABUNIBORDERS struct {
	Name struct {
		Name string `json:"common"`
	} `json:"name"`
	Borders []string `json:"borders"`
}
