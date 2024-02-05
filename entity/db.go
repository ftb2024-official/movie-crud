package entity

type Director struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
}

type Movie struct {
	Id       string   `json:"id,omitempty"`
	Isbn     string   `json:"isbn"`
	Title    string   `json:"title"`
	Year     int      `json:"year"`
	Director Director `json:"director"`
}
