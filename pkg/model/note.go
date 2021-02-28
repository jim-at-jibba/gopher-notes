package model

type NewNote struct {
	Text   string `json:"text"`
	Title  string `json:"title"`
	UserID string `json:"userId"`
}

type Note struct {
	ID    string `json:"id"`
	Text  string `json:"text"`
	Title string `json:"title"`
	User  *User  `json:"user"`
}
