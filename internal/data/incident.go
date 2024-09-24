package data

import (
	"context"
	"github.com/tang95/x-seek/internal/model"
)

type incidentRepo struct {
	*Data
}

func newIncidentRepo(data *Data) model.IncidentRepo {
	return &incidentRepo{data}
}

func (repo *incidentRepo) Get(ctx context.Context, id string) (*model.Incident, error) {
	component := model.Incident{}
	tx := repo.DB(ctx).Where("id = ?", id).First(&component)
	return &component, tx.Error
}

func (repo *incidentRepo) Create(ctx context.Context, incident *model.Incident) (string, error) {
	tx := repo.DB(ctx).Create(incident)
	return incident.ID, tx.Error
}

func (repo *incidentRepo) Update(ctx context.Context, id string, incident *model.Incident) error {
	return repo.DB(ctx).Model(&model.Incident{}).Where("id = ?", id).Updates(incident).Error
}

func (repo *incidentRepo) Query(ctx context.Context, filter *model.IncidentFilter, page *model.PageQuery, sort []*model.SortQuery) ([]*model.Incident, int64, error) {
	var (
		incidents []*model.Incident
		total     int64
	)
	tx := repo.DB(ctx).Model(&model.Incident{})
	if filter.Keywords != "" {
		tx = tx.Where("name like ?", "%"+filter.Keywords+"%")
	}
	if filter.Status != "" {
		tx = tx.Where("status = ?", filter.Status)
	}
	if filter.Severity != "" {
		tx = tx.Where("severity = ?", filter.Severity)
	}
	tx = tx.Count(&total)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	if page != nil {
		tx = tx.Offset(page.GetOffset()).Limit(page.GetLimit())
	}
	if sort != nil {
		for _, s := range sort {
			tx = tx.Order(s.Field + " " + string(s.Direction))
		}
	}
	tx = tx.Find(&incidents)
	return incidents, total, tx.Error
}

func (repo *incidentRepo) Count(ctx context.Context, filter *model.IncidentFilter) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *incidentRepo) AddUser(ctx context.Context, incidentUser *model.IncidentUser) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *incidentRepo) RemoveUser(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (repo *incidentRepo) QueryUser(ctx context.Context, incidentID string, filter *model.IncidentUserFilter, page *model.PageQuery, sort []*model.SortQuery) ([]*model.IncidentUser, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *incidentRepo) AddActivity(ctx context.Context, activity model.IncidentActivity[any]) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *incidentRepo) RemoveActivity(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (repo *incidentRepo) QueryActivity(ctx context.Context, incidentID string, filter *model.IncidentActivityFilter, page *model.PageQuery, sort []*model.SortQuery) ([]*model.IncidentActivity[any], int64, error) {
	//TODO implement me
	panic("implement me")
}
