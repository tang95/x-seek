package model

import "time"

type IncidentActivityType string

const (
	Comment     IncidentActivityType = "comment"
	AddUser     IncidentActivityType = "add_user"
	RemoveUser  IncidentActivityType = "remove_user"
	SetSeverity IncidentActivityType = "set_severity"
	SetStatus   IncidentActivityType = "set_status"
)

type IncidentActivity[T any] struct {
	BaseModel
	IncidentID string               `gorm:"type:varchar(36); not null; comment:事件ID" json:"incident_id"`
	UserID     string               `gorm:"type:varchar(36); not null; comment:用户ID" json:"user_id"`
	HappenedAt time.Time            `gorm:"type:datetime; not null; comment:发生时间" json:"happened_at"`
	Type       IncidentActivityType `gorm:"type:varchar(20); not null; comment:类型" json:"type"`
	Data       T                    `gorm:"type:json; not null; serializer:json; comment:数据" json:"data"`
}

func (IncidentActivity[any]) TableName() string {
	return "incident_activity"
}

type IncidentActivityComment struct {
	Text string `json:"text"`
}

type IncidentActivityUser struct {
	User *User `json:"user"`
}

type IncidentActivitySetSeverity struct {
	Source Severity `json:"source"`
	Target Severity `json:"target"`
}

type IncidentActivitySetStatus struct {
	Source IncidentStatus `json:"source"`
	Target IncidentStatus `json:"target"`
}
