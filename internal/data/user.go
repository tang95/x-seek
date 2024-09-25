package data

import (
	"context"
	"errors"
	"github.com/tang95/x-seek/internal/model"
	"gorm.io/gorm"
)

type userRepo struct {
	*Data
}

func newUserRepo(data *Data) model.UserRepo {
	return &userRepo{data}
}

func (repo *userRepo) Count(ctx context.Context) (int64, error) {
	var count int64
	tx := repo.DB(ctx).Model(&model.User{}).Count(&count)
	return count, tx.Error
}

func (repo *userRepo) GetByDingtalkID(ctx context.Context, unionId string) (*model.User, error) {
	users := make([]*model.User, 0)
	tx := repo.DB(ctx).Find(&users, "dingtalk_id = ?", unionId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if len(users) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	if len(users) > 1 {
		return nil, errors.New("found more than one user")
	}
	return users[0], nil
}

func (repo *userRepo) GetByGithubID(ctx context.Context, githubID string) (*model.User, error) {
	users := make([]*model.User, 0)
	tx := repo.DB(ctx).Find(&users, "github_id = ?", githubID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if len(users) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	if len(users) > 1 {
		return nil, errors.New("found more than one user")
	}
	return users[0], nil
}

func (repo *userRepo) Query(ctx context.Context, filter *model.UserFilter, page *model.PageQuery, sort []*model.SortQuery) ([]*model.User, int64, error) {
	var (
		users []*model.User
		total int64
	)
	tx := repo.DB(ctx).Model(&model.User{})
	tx = tx.Count(&total)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	if page != nil {
		tx = tx.Offset(page.GetOffset()).
			Limit(page.GetLimit())
	}
	if sort != nil {
		for _, s := range sort {
			tx = tx.Order(s.Field + " " + string(s.Direction))
		}
	}
	tx = tx.Find(&users)
	return users, total, tx.Error
}

func (repo *userRepo) Create(ctx context.Context, user *model.User) (*model.User, error) {
	tx := repo.DB(ctx).Create(user)
	return user, tx.Error
}

func (repo *userRepo) Get(ctx context.Context, id string) (*model.User, error) {
	user := model.User{}
	tx := repo.DB(ctx).Where("id = ?", id).First(&user)
	return &user, tx.Error
}

func (repo *userRepo) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (repo *userRepo) Update(ctx context.Context, id string, user *model.User) error {
	//TODO implement me
	panic("implement me")
}
