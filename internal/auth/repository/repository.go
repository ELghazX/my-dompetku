package repository

import (
	"context"

	"github.com/elghazx/my-dompetku/internal/auth/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdatePassword(ctx context.Context, id uint, password string) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(ctx context.Context, user *domain.User) error {
	return ur.db.WithContext(ctx).Create(user).Error
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := ur.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) UpdatePassword(ctx context.Context, id uint, password string) error {
	return ur.db.WithContext(ctx).Model(&domain.User{}).
		Where("id = ?", id).
		Update("password", password).Error
}
