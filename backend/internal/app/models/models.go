package models

type Track struct {
	ID     string `json:"value"`
	Number string `json:"text"`
	Driver string
} 

type Point struct {
	Lable    string   `json:"label"`
	Title    string   `json:"title"`
	Position Position `json:"position"`
}

type Position struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// var DefaultPoint Point = Point{Lat: 47.012271881103516, Lng: 28.860593795776367}
