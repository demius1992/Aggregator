package models

type Areas struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	MapMarker   string `json:"mapMarker"`
	Price       string `json:"price"`
	Info        string `json:"-"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Typ         string `json:"typ"`
}

type GetAllAreas struct {
	Data []Areas `json:"data"`
}
