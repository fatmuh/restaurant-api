package response

type LoginResponse struct {
	ExpiresIn string `json:"expires_in"`
	Token     string `json:"token"`
	IsActive  string `json:"is_active"`
}
