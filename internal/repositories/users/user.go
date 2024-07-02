package users

import (
	"context"

	"github.com/arieffian/go-boilerplate/internal/models"
	"github.com/arieffian/providers/pkg/redis"
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

func (r *userRepository) GetUserById(ctx context.Context, p GetUserByIdParams) (*GetUserByIdResponse, error) {
	var user models.User

	err := r.redisDb.GetCache(ctx, p.UserId, &user)

	if err != nil {
		if err != redis_pkg.Nil {
			return nil, err
		}

		if err := r.db.First(&user, "id = ?", p.UserId).Error; err != nil {
			return nil, err
		}
		if err := r.redisDb.SetCacheWithExpiration(context.Background(), p.UserId, user, 0); err != nil {
			return nil, err
		}
	}

	return &GetUserByIdResponse{
		User: user,
	}, nil
}

func NewUserRepository(p NewUserRepositoryParams) *userRepository {

	return &userRepository{
		db:      p.Db,
		redisDb: p.Redis,
	}
}
