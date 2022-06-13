package models

type Flats struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	MapMarker   string `json:"mapMarker"`
	Price       string `json:"price"`
	Info        string `json:"info"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Rooms       int    `json:"rooms"`
	Typ         string `json:"typ"`
}

type GetAllFlats struct {
	Data []Flats `json:"data"`
}
