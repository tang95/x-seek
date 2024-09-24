package model

import (
	"context"
	"time"
)

type IncidentRole string

const (
	Commander IncidentRole = "commander"
	Responder IncidentRole = "responder"
)

type IncidentStatus string

const (
	Active   IncidentStatus = "active"
	Resolved IncidentStatus = "resolved"
)

type Severity string

const (
	// Critical 严重问题，需公开通知和高层团队介入，影响大量客户
	Critical Severity = "critical"
	// High 重要问题，影响大量客户，需快速处理
	High Severity = "high"
	// Medium 一般问题，需立即关注，有升级为更严重问题的可能
	Medium Severity = "medium"
	// Low 轻微问题，不影响客户使用产品功能
	Low Severity = "low"
)

type IncidentUser struct {
	BaseModel
	UserID     string       `gorm:"type:varchar(36); not null; comment:用户ID" json:"user_id"`
	Role       IncidentRole `gorm:"type:varchar(20); not null; comment:角色" json:"role"`
	IncidentID string       `gorm:"type:varchar(36); not null; comment:事件ID" json:"incident_id"`
}

func (IncidentUser) TableName() string {
	return "incident_user"
}

type Incident struct {
	BaseModel
	Name        string               `gorm:"type:varchar(100); not null; comment:名称" json:"name"`
	Description string               `gorm:"type:text; comment:描述" json:"description"`
	Status      IncidentStatus       `gorm:"type:varchar(20); not null; comment:状态" json:"status"`
	Severity    Severity             `gorm:"type:varchar(20); not null; comment:严重程度" json:"severity"`
	Tags        []string             `gorm:"type:json; comment:标签" json:"tags"`
	ReportedAt  time.Time            `gorm:"type:datetime; not null; comment:报告时间" json:"reported_at"`
	ReporterID  string               `gorm:"type:varchar(36); not null; comment:报告人ID" json:"reporter_id"`
	ResolvedAt  *time.Time           `gorm:"type:datetime; comment:解决时间" json:"resolved_at"`
	Channel     IncidentChannel[any] `gorm:"type:json; comment:沟通渠道" json:"channel"`
}

func (Incident) TableName() string {
	return "incident"
}

type IncidentFilter struct {
	Keywords string
	Status   IncidentStatus
	Severity Severity
}

type IncidentUserFilter struct {
}

type IncidentActivityFilter struct {
}

type IncidentRepo interface {
	Get(ctx context.Context, id string) (*Incident, error)
	Create(ctx context.Context, incident *Incident) (string, error)
	Update(ctx context.Context, id string, incident *Incident) error
	Query(ctx context.Context, filter *IncidentFilter, page *PageQuery, sort []*SortQuery) ([]*Incident, int64, error)
	AddUser(ctx context.Context, incidentUser *IncidentUser) (string, error)
	RemoveUser(ctx context.Context, id string) error
	QueryUser(ctx context.Context, incidentID string, filter *IncidentUserFilter, page *PageQuery, sort []*SortQuery) ([]*IncidentUser, int64, error)
	AddActivity(ctx context.Context, activity IncidentActivity[any]) (string, error)
	RemoveActivity(ctx context.Context, id string) error
	QueryActivity(ctx context.Context, incidentID string, filter *IncidentActivityFilter, page *PageQuery, sort []*SortQuery) ([]*IncidentActivity[any], int64, error)
}
