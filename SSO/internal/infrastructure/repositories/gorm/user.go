package gormrepository

import (
	"SSO/internal/domain"
	"SSO/internal/infrastructure/repositories/gorm/entities"
	"SSO/internal/usecases"
	"SSO/internal/utils/mappers"
	"context"
	"errors"
	"gorm.io/gorm"
	"log/slog"
)

var _ usecases.UserRepositoryInterface = &GormUserRepository{}

// GormUserRepository is a repository for Gorm databases.
type GormUserRepository struct {
	AbstractGormRepository[*domain.User, entities.User]
}

// NewGormUserRepository creates a new GormUserRepository.
func NewGormUserRepository(db *gorm.DB, logger *slog.Logger) *GormUserRepository {
	return &GormUserRepository{
		AbstractGormRepository: NewAbstractGormRepository[*domain.User, entities.User](
			db,
			logger,
			mappers.UserDomainToUserEntity,
			mappers.UserEntityToUserDomain,
		),
	}
}

// GetByUsername gets a user by username.
func (r *GormUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	const op = "GormUserRepository.GetByUsername"
	var entity entities.User
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&entity).Error; err != nil {
		r.logger.Error(op, slog.Any("error", err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return r.entityToModel(&entity), nil
}

// GetByEmail gets a user by email.
func (r *GormUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	const op = "GormUserRepository.GetByEmail"
	var entity entities.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&entity).Error; err != nil {
		r.logger.Error(op, slog.Any("error", err.Error()))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return r.entityToModel(&entity), nil
}
