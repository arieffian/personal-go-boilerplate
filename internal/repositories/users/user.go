package users

import (
	"context"
	"fmt"

	"github.com/arieffian/go-boilerplate/internal/config"
	"github.com/arieffian/go-boilerplate/internal/models"
	database "github.com/arieffian/providers/pkg/db"
	"github.com/arieffian/providers/pkg/redis"
	redis_pkg "github.com/redis/go-redis/v9"
)

const (
	GetUserByIdCacheKey = "|%s|get-user-by-id|id|%s|"
	GetUsersCacheKey    = "|%s|get-users|offset|%d|limit|%d|"
)

type userRepository struct {
	db      *database.DbInstance
	redisDb redis.RedisService
	cfg     *config.Config
}

var _ UserRepository = (*userRepository)(nil)

type NewUserRepositoryParams struct {
	Db    *database.DbInstance
	Redis redis.RedisService
	Cfg   *config.Config
}

func (r *userRepository) GetUserById(ctx context.Context, p GetUserByIdParams) (*GetUserByIdResponse, error) {
	var user models.User

	cacheKey := fmt.Sprintf(GetUserByIdCacheKey, r.cfg.ServiceName, p.UserId)

	err := r.redisDb.GetCache(ctx, cacheKey, &user)

	if err != nil {
		if err != redis_pkg.Nil {
			return nil, err
		}

		if err := r.db.Db.First(&user, "id = ?", p.UserId).Error; err != nil {
			return nil, err
		}
		if err := r.redisDb.SetCacheWithExpiration(context.Background(), cacheKey, user, r.cfg.CacheTTL); err != nil {
			return nil, err
		}
	}

	return &GetUserByIdResponse{
		User: user,
	}, nil
}

func (r *userRepository) GetUsers(ctx context.Context, p GetUsersParams) (*GetUsersResponse, error) {

	var users []models.User

	cacheKey := fmt.Sprintf(GetUsersCacheKey, r.cfg.ServiceName, p.Offset, p.Limit)

	err := r.redisDb.GetCache(ctx, cacheKey, &users)

	if err != nil {
		if err != redis_pkg.Nil {
			return nil, err
		}

		err := r.db.Db.Model(&models.User{}).Limit(p.Limit).Offset(p.Offset).Find(&users).Error

		if err != nil {
			return nil, err
		}

		if err := r.redisDb.SetCacheWithExpiration(context.Background(), cacheKey, users, r.cfg.CacheTTL); err != nil {
			return nil, err
		}
	}

	return &GetUsersResponse{
		Users: users,
	}, nil

}

func (r *userRepository) CreateNewUser(ctx context.Context, p CreateNewUserParams) (*CreateNewUserResponse, error) {
	user := models.User{
		Name: p.Name,
	}

	err := r.db.Db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &CreateNewUserResponse{
		User: user,
	}, nil
}

func NewUserRepository(p NewUserRepositoryParams) UserRepository {

	return &userRepository{
		db:      p.Db,
		redisDb: p.Redis,
		cfg:     p.Cfg,
	}
}
