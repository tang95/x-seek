package data

import (
	"context"
	"github.com/tang95/x-seek/internal/model"
)

type teamRepo struct {
	*Data
}

func newTeamRepo(data *Data) model.TeamRepo {
	return &teamRepo{data}
}

func (repo *teamRepo) Create(ctx context.Context, team *model.Team) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *teamRepo) Get(ctx context.Context, id string) (*model.Team, error) {
	team := &model.Team{}
	tx := repo.DB(ctx).Where("id = ?", id).First(&team)
	return team, tx.Error
}

func (repo *teamRepo) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (repo *teamRepo) Update(ctx context.Context, id string, team *model.Team) error {
	//TODO implement me
	panic("implement me")
}

func (repo *teamRepo) Query(ctx context.Context, filter *model.TeamFilter, page *model.PageQuery, sort []*model.SortQuery) ([]*model.Team, int64, error) {
	var (
		teams []*model.Team
		total int64
	)
	tx := repo.DB(ctx).Model(&model.Team{})
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
	tx = tx.Find(&teams)
	return teams, total, tx.Error
}

func (repo *teamRepo) QueryMember(ctx context.Context, id string, filter *model.UserFilter, page *model.PageQuery, sort []*model.SortQuery) ([]*model.User, int64, error) {
	var (
		users []*model.User
		total int64
	)
	tx := repo.DB(ctx).Model(&model.User{}).
		Joins("inner join team_user on user_id = user.id and team_id = ?", id)
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

func (repo *teamRepo) Count(ctx context.Context, filter *model.TeamFilter) (int64, error) {
	var (
		total int64
	)
	tx := repo.DB(ctx).Model(&model.Team{}).Count(&total)
	return total, tx.Error
}
