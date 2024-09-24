package model

import "context"

type Team struct {
	BaseModel
	Name        string `gorm:"type:varchar(100); not null; comment:名称"`
	OwnerID     string `gorm:"type:varchar(36); not null; comment:拥有者ID"`
	Description string `gorm:"type:text; comment:描述"`
}

func (Team) TableName() string {
	return "team"
}

type TeamUser struct {
	BaseModel
	TeamID string `gorm:"type:varchar(36); not null; comment:团队ID"`
	UserID string `gorm:"type:varchar(36); not null; comment:用户ID"`
}

func (TeamUser) TableName() string {
	return "team_user"
}

type TeamFilter struct {
}

type TeamRepo interface {
	Create(ctx context.Context, team *Team) (string, error)
	Get(ctx context.Context, id string) (*Team, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, team *Team) error
	Query(ctx context.Context, filter *TeamFilter, page *PageQuery, sort []*SortQuery) ([]*Team, int64, error)
	QueryMember(ctx context.Context, id string, filter *UserFilter, page *PageQuery, sort []*SortQuery) ([]*User, int64, error)
	Count(ctx context.Context, filter *TeamFilter) (int64, error)
}
