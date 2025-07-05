package service

import (
	"context"

	"github.com/elghazx/my-dompetku/internal/auth/repository"
)

type authService struct {
	userRepository repository.UserRepository
}

type AuthService interface {
	Login(ctx context.Context, username,  password string) (dto.auth)
}
