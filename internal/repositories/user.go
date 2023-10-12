package repositories

import (
	"github.com/arieffian/go-boilerplate/internal/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type userRepository struct {
	db      *gorm.DB
	redisDb *redis.Client
}

var _ UserInterface = (*userRepository)(nil)

type NewUserRepositoryParams struct {
	Db    *gorm.DB
	Redis *redis.Client
}

func (r *userRepository) GetUserById(id string) (*models.User, error) {
	var user models.User

	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(p NewUserRepositoryParams) *userRepository {

	return &userRepository{
		db:      p.Db,
		redisDb: p.Redis,
	}
}
