package models

type Flats struct {
	Id          int32  `json:"-"`
	Title       string `json:"title"`
	MapMarker   string `json:"mapMarker"`
	Price       string `json:"price"`
	Info        string `json:"info"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Rooms       int32  `json:"rooms"`
	Type        string `json:"type"`
}

type GetAllFlats struct {
	Data []Flats `json:"data"`
}
