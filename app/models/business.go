package models

type Businesses struct {
	Categories []struct {
		Alias string `json:"alias"`
		Title string `json:"title"`
	} `json:"categories"`

	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`

	DisplayPhone string  `json:"display_phone"`
	Distance     float64 `json:"distance"`
	ID           string  `json:"id"`
	ImageURL     string  `json:"image_url"`
	IsClosed     bool    `json:"is_closed"`
	Location     struct {
			   Address1       string   `json:"address1"`
			   Address2       string   `json:"address2"`
			   Address3       string   `json:"address3"`
			   City           string   `json:"city"`
			   Country        string   `json:"country"`
			   DisplayAddress []string `json:"display_address"`
			   State          string   `json:"state"`
			   ZipCode        string   `json:"zip_code"`
		   } `json:"location"`

	Name         string        `json:"name"`
	Phone        string        `json:"phone"`
	Price        string        `json:"price"`
	Rating       float64       `json:"rating"`
	ReviewCount  int           `json:"review_count"`
	Transactions []interface{} `json:"transactions"`
	URL          string        `json:"url"`
}