package service

import (
	"roastkuy-api/data/request"
	"time"
)

type AuthService interface {
	Login(account request.LoginRequest) (string, time.Time, error)
	Register(account request.CreateAccountRequest)
}
