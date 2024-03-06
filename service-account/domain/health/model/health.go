package model

type HealthCheck struct {
	AppDetail      AppDetail      `json:"app"`
	DatabaseDetail DatabaseDetail `json:"database"`
	Status         string         `json:"status"`
}

type AppDetail struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type DatabaseDetail struct {
	Dialect string `json:"driver"`
	Status  string `json:"status"`
}
