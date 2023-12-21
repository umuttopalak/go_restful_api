package todos

type todo struct {
	ID          int    `json:"id"`
	Author      string `json:"author"`
	Description string `json:"desc"`
}
