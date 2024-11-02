package domain

type Sample struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Age       int     `json:"age"`
	Size      float64 `json:"size"`
	IsVisible bool    `json:"is_visible"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
