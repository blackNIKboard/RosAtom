package sensor

import (
	"container/list"
	"math/rand"
	"rosatomcase/backend/mock"
	"rosatomcase/backend/model"
	"time"
)

type Array struct {
	Array []*Sensor
}

func (receiver Array) Retrieve() []model.SensorData {
	var data []model.SensorData

	for i := 0; i < len(receiver.Array); i++ {
		data = append(data, receiver.Array[i].Last())
	}

	return data
}

type Sensor struct {
	sensors list.List
}

func tempHealthyData() int {
	return rand.Int()%10 + 20
}

func tempUnhealthyData() int {
	return rand.Int()%21 + 40
}

func brightHealthyData() int {
	return rand.Int()%400 + 100
}

func brightUnhealthyData() int {
	return rand.Int() % 1000
}

func energyHealthyData() int {
	return rand.Int() % 1000
}

func energyUnhealthyData() int {
	return rand.Int()%1000 + 1000
}

func (receiver *Sensor) checkValue(val float32, warn model.ValueWarning) bool {
	if val > warn.Maximum || val < warn.Minimum {
		return false
	}

	return true
}

func (receiver *Sensor) Generate(name string, tempWarn model.ValueWarning, energyWarn model.ValueWarning) {
	errProb := 0.99

	randTemp := mock.Randomizer{
		Value:       tempHealthyData,
		Unvalue:     tempUnhealthyData,
		ProbOfValue: errProb,
	}

	randBright := mock.Randomizer{
		Value:       brightHealthyData,
		Unvalue:     brightUnhealthyData,
		ProbOfValue: errProb,
	}

	randEnergy := mock.Randomizer{
		Value:       energyHealthyData,
		Unvalue:     energyUnhealthyData,
		ProbOfValue: errProb,
	}

	recentErr := 10
	recentErrTempEx := false
	recentErrBrightEx := false
	recentErrEnergyEx := false

	for i := 0; ; i++ {
		temp := float32(randTemp.Return().(func() int)())
		bright := float32(randBright.Return().(func() int)())
		energy := float32(randEnergy.Return().(func() int)())

		tempHealth := receiver.checkValue(temp, tempWarn)
		brightHealth := receiver.checkValue(bright, model.ValueWarning{
			Minimum: 100,
			Maximum: 500,
		})
		energyHealth := receiver.checkValue(energy, energyWarn)
		health := tempHealth && brightHealth && energyHealth

		if health == false {
			recentErr = 0
			recentErrTempEx = tempHealth
			recentErrBrightEx = brightHealth
			recentErrEnergyEx = energyHealth
		} else {
			recentErr++
		}

		receiver.sensors.PushBack(model.SensorData{
			Name:   name,
			Mapped: "sda21321",
			Values: model.SensorValues{
				Temperature: temp,
				Brightness:  bright,
				Energy:      energy,
			},
			Health: model.HealthCheck{
				Health: func() bool {
					if recentErr > 5 {
						return true
					} else {
						return false
					}
				}(),
				Temperature: func() bool {
					if recentErr > 5 {
						return tempHealth
					} else {
						return recentErrTempEx
					}
				}(),
				Brightness: func() bool {
					if recentErr > 5 {
						return brightHealth
					} else {
						return recentErrBrightEx
					}
				}(),
				Energy: func() bool {
					if recentErr > 5 {
						return energyHealth
					} else {
						return recentErrEnergyEx
					}
				}(),
			},
		})

		time.Sleep(time.Second / 2)

		for i := 0; i < receiver.sensors.Len()-10; i++ {
			receiver.sensors.Remove(receiver.sensors.Front())
		}
	}
}

func (receiver *Sensor) Last() model.SensorData {
	//spew.Dump(receiver.sensors)
	return receiver.sensors.Back().Value.(model.SensorData)
}
