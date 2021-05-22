package sensor

import (
	"container/list"
	"github.com/davecgh/go-spew/spew"
	wr "github.com/mroth/weightedrand"
	"math/rand"
	"rosatomcase/backend/model"
	"time"
)

type Array struct {
	Array []*Sensor
}

func (receiver Array) Retrieve() []model.SensorData {
	var data []model.SensorData
	spew.Dump("retrieving data")

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
	chooserTemp, _ := wr.NewChooser(
		wr.Choice{Item: tempHealthyData, Weight: 9},
		wr.Choice{Item: tempUnhealthyData, Weight: 1},
	)

	chooserBright, _ := wr.NewChooser(
		wr.Choice{Item: brightHealthyData, Weight: 9},
		wr.Choice{Item: brightUnhealthyData, Weight: 1},
	)

	chooserEnergy, _ := wr.NewChooser(
		wr.Choice{Item: energyHealthyData, Weight: 9},
		wr.Choice{Item: energyUnhealthyData, Weight: 1},
	)

	recent_err := 10

	for i := 0; ; i++ {
		temp := float32(chooserTemp.Pick().(func() int)())
		bright := float32(chooserBright.Pick().(func() int)())
		energy := float32(chooserEnergy.Pick().(func() int)())

		tempHealth := receiver.checkValue(temp, tempWarn)
		brightHealth := receiver.checkValue(bright, model.ValueWarning{
			Minimum: 100,
			Maximum: 500,
		})
		energyHealth := receiver.checkValue(energy, energyWarn)
		health := tempHealth && brightHealth && energyHealth

		if health == false {
			recent_err = 0
		} else {
			recent_err++
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
					if recent_err > 5 {
						return true
					} else {
						return false
					}
				}(),
				Temperature: tempHealth,
				Brightness:  brightHealth,
				Energy:      energyHealth,
			},
		})

		time.Sleep(time.Second)

		for i := 0; i < receiver.sensors.Len()-10; i++ {
			receiver.sensors.Remove(receiver.sensors.Front())
		}
	}
}

func (receiver *Sensor) Last() model.SensorData {
	//spew.Dump(receiver.sensors)
	return receiver.sensors.Back().Value.(model.SensorData)
}
