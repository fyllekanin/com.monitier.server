package entity

type PingEntity struct {
	ServiceName  string `json:"serviceName"`
	ResponseTime int    `json:"responseTime"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
}
