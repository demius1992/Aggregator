package models

type Cottages struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	MapMarker   string `json:"mapMarker"`
	Price       string `json:"price"`
	Info        string `json:"info"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Rooms       int    `json:"rooms"`
	Type        string `json:"type"`
}

type GetAllCottages struct {
	Data []Cottages `json:"data"`
}