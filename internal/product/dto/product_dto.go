package dto


type Product struct{
	ID uint64 `json:"id"`
	Image string `json:"image"`
	Title string `json:"title"`
	Price float64 `json:"price"`
	Description string `json:"description"`
}