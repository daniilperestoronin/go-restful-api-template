package record

type Record struct {
	Id    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
}
