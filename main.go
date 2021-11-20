package main

import (
	"time"

	"github.com/shuaiqidechuan/co2-raspberry/exporter"
	"github.com/shuaiqidechuan/co2-raspberry/operator"
	"github.com/shuaiqidechuan/co2-raspberry/util/log"

	sensor "github.com/shuaiqidechuan/co2-raspberry/co2_sensor"
	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/serial0", Baud: 115200, ReadTimeout: time.Second * 5}
	co2Sensor, err := sensor.Connect(c)
	if err != nil {
		log.Fatal(err)
	}

	co2Operator := operator.NewOperator(co2Sensor)

	e := exporter.NewExporter()
	e.Register("co2", func() (interface{}, error) {
		return co2Operator.QueryCO2()
	})

	if err := e.Run(); err != nil {
		log.Fatal(err)
	}
}
