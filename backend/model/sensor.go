package model

type SensorData struct {
	Name   string       `json:"name"`
	Mapped string       `json:"mapped"`
	Values SensorValues `json:"values"`
	Health HealthCheck  `json:"health"`
}

type SensorValues struct {
	Temperature float32 `json:"temperature"`
	Brightness  float32 `json:"brightness"`
	Energy      float32 `json:"energy"`
}

type ValueWarning struct {
	Minimum float32
	Maximum float32
}

type HealthCheck struct {
	Health      bool `json:"health"`
	Temperature bool `json:"temperature"`
	Brightness  bool `json:"brightness"`
	Energy      bool `json:"energy"`
}
