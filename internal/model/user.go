package model

import "context"

type Role string

const (
	Admin  Role = "admin"
	Member Role = "member"
)

type User struct {
	BaseModel
	Name        string `gorm:"type:varchar(100); not null; comment:名称"`
	Description string `gorm:"type:text; comment:描述"`
	Avatar      string `gorm:"type: varchar(255); comment:头像"`
	Role        Role   `gorm:"type:varchar(20); not null; comment:角色"`
	DingTalkID  string `gorm:"type:varchar(100); comment:钉钉ID"`
}

func (User) TableName() string {
	return "user"
}

type UserFilter struct {
}

type UserRepo interface {
	Query(ctx context.Context, filter *UserFilter, page *PageQuery, sort []*SortQuery) ([]*User, int64, error)
	Create(ctx context.Context, user *User) (*User, error)
	Get(ctx context.Context, id string) (*User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, user *User) error
	GetByGithubID(ctx context.Context, githubID string) (*User, error)
	Count(ctx context.Context) (int64, error)
}
