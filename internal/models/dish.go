package models

type Dish struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Image       string  `json:"image" db:"image"`
	VideoLink   string  `json:"video_link" db:"video_link"`
	Calories    float64 `json:"calories" db:"calories"`
	Variation   int     `json:"variation" db:"variation"`
	DayTime     int     `json:"day_time" db:"day_time"`
}
