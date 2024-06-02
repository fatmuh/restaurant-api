package request

type CreateAccountRequest struct {
	Name     string `validate:"required,min=3,max=100" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6,max=100" json:"password"`
	Phone    string `validate:"required,min=11,max=13" json:"phone"`
}

type UpdateAccountRequest struct {
	Id    int    `validate:"required,min=1" json:"id"`
	Name  string `validate:"required,min=3,max=100" json:"name"`
	Phone string `validate:"required,min=11,max=13" json:"phone"`
}

type LoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6,max=100" json:"password"`
}
