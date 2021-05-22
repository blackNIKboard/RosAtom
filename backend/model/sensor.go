package model

type SensorData struct {
	Name   string  `json:"name"`
	Mapped string  `json:"mapped"`
	Value  float32 `json:"value"`
	Health bool    `json:"health"`
}
