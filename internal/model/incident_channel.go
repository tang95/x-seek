package model

type IncidentChannelType string

const (
	DingTalk IncidentChannelType = "dingtalk"
	Feishu   IncidentChannelType = "feishu"
)

type IncidentChannel[T any] struct {
	Data T                   `json:"data"`
	Type IncidentChannelType `json:"type"`
}

type IncidentChannelDingTalk struct {
	Link string `json:"link"`
}

type IncidentChannelFeishu struct {
	Link string `json:"link"`
}
