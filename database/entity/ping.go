package entity

type PingEntity struct {
	ServiceName  string `json:"serviceName"`
	ResponseTime int    `json:"responseTime"`
	CreatedAt    int    `json:"createdAt"`
	UpdatedAt    int    `json:"updatedAt"`
}
