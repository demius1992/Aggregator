package models

type Areas struct {
	Id          int    `json:"-"`
	Title       string `json:"title,omitempty"`
	MapMarker   string `json:"mapMarker,omitempty"`
	Price       string `json:"price,omitempty"`
	Info        string `json:"-,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
	Typ         string `json:"typ,omitempty"`
}

type GetAllAreas struct {
	Data []Areas `json:"data"`
}
