package response

type MenusResponse struct {
	Id          int     `json:"id"`
	MenuName    string  `json:"menu_name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Price       int     `json:"price"`
	Discount    float32 `json:"discount"`
}
