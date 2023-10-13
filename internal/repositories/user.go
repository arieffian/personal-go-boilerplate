package repositories

import (
	"context"

	"github.com/arieffian/go-boilerplate/internal/models"
	"github.com/arieffian/go-boilerplate/internal/pkg/redis"
	redis_pkg "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type userRepository struct {
	db      *gorm.DB
	redisDb redis.RedisService
}

var _ UserInterface = (*userRepository)(nil)

type NewUserRepositoryParams struct {
	Db    *gorm.DB
	Redis redis.RedisService
}

// todo: add ctx on every function
func (r *userRepository) GetUserById(id string) (*models.User, error) {
	var user models.User

	err := r.redisDb.GetCache(context.Background(), id, &user)

	if err != nil {
		if err != redis_pkg.Nil {
			return nil, err
		}

		if err := r.db.First(&user, "id = ?", id).Error; err != nil {
			return nil, err
		}
		if err := r.redisDb.SetCacheWithExpiration(context.Background(), id, user, 0); err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func NewUserRepository(p NewUserRepositoryParams) *userRepository {

	return &userRepository{
		db:      p.Db,
		redisDb: p.Redis,
	}
}
