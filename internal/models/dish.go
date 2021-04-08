package models

type Dish struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	VideoLink   string  `json:"video_link"`
	Calories    float64 `json:"calories"`
	Variation   int     `json:"variation"`
	DayTime     int     `json:"day_time"`
}
