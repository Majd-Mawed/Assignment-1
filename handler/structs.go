package handler

type UNI struct {
	Name     string   `json:"name"`
	Country  string   `json:"country"`
	Webpages []string `json:"web_pages"`
	Isocode  string   `json:"alpha_two_code"`
}

type NABUNI struct {
	Name      string            `json:"name"`
	Isocode   string            `json:"alpha_two_code"`
	languages map[string]string `json:"languages"`
}
