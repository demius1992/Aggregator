package models

type Cottages struct {
	Id          int    `json:"-,omitempty"`
	Title       string `json:"title,omitempty"`
	MapMarker   string `json:"mapMarker,omitempty"`
	Price       string `json:"price,omitempty"`
	Info        string `json:"info,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
	Rooms       int    `json:"rooms,omitempty"`
	Typ         string `json:"typ,omitempty"`
}

type GetAllCottages struct {
	Data []Cottages `json:"data"`
}
