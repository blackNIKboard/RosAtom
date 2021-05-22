package model

type Sensor struct {
	Name   string  `json:"name"`
	Mapped string  `json:"mapped"`
	Value  float32 `json:"value"`
	Health bool    `json:"health"`
}
