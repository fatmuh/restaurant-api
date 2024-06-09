package response

import "time"

type PromosResponse struct {
	PromoName       string    `json:"promo_name"`
	Image           string    `json:"image"`
	Type            string    `json:"type"`
	OutdatePromo    time.Time `json:"outdate"`
	Description     string    `json:"description"`
	DetailTutorial  string    `json:"detail_tutorial"`
	DetailCondition string    `json:"detail_condition"`
}
