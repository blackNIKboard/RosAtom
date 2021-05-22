package temperature

import (
	"container/list"
	"github.com/davecgh/go-spew/spew"
	wr "github.com/mroth/weightedrand"
	"math"
	"math/rand"
	"rosatomcase/backend/model"
	"time"
)

type TempArray struct {
	Array map[string]*Temperature
}

type Temperature struct {
	sensors list.List
}

func healthyData() int {
	return rand.Int()%10 + 20
}

func unhealthyData() int {
	return rand.Int()%21 + 40
}

func (receiver *Temperature) Generate(name string, maxRange float64) {
	chooser, _ := wr.NewChooser(
		wr.Choice{Item: healthyData, Weight: 9},
		wr.Choice{Item: unhealthyData, Weight: 1},
	)

	for i := 0; ; i++ {
		gen := chooser.Pick().(func() int)
		val := float32(gen())
		health := true

		if i > 0 && math.Abs(float64(val-receiver.sensors.Back().Value.(model.SensorData).Value)) > maxRange {
			health = false
		}

		receiver.sensors.PushBack(model.SensorData{
			Name:   name,
			Mapped: "sda21321",
			Value:  val,
			Health: health,
		})

		time.Sleep(time.Second)

		for i := 0; i < receiver.sensors.Len()-10; i++ {
			receiver.sensors.Remove(receiver.sensors.Front())
		}
	}
}

func (receiver *Temperature) Last() model.SensorData {
	spew.Dump(receiver.sensors)

	return receiver.sensors.Back().Value.(model.SensorData)
}
