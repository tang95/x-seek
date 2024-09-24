package service

import (
	"context"
	"github.com/patrickmn/go-cache"
	"github.com/tang95/x-seek/internal/model"
)

const (
	CountUserCacheKey = "count_user"
)

func (service *Service) GetUserByGithubID(ctx context.Context, githubID string) (*model.User, error) {
	user, err := service.userRepo.GetByGithubID(ctx, githubID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *Service) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	// 如果是第一个用户，则默认为管理员
	count, ok := service.cache.Get(CountUserCacheKey)
	if !ok {
		count, _ = service.userRepo.Count(ctx)
	}
	if count.(int32) == 0 {
		user.Role = model.Admin
	}
	newUser, err := service.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	service.cache.Set(CountUserCacheKey, int32(1), cache.NoExpiration)
	return newUser, nil
}
