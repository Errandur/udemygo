package main

import (
	"fmt"
)

const u16BitMax float64 = 65535
const kmhMultiple float64 = 1.60934

type Car struct {
	gasPedal      uint16 // 0-65535
	brakePedal    uint16
	steeringWheel int16 // (-32k) - (+32k)
	topSpeedKMH   float64
}

func (c Car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKMH / u16BitMax)
}

func (c Car) milesH() float64 {
	return float64(c.gasPedal) * (c.topSpeedKMH / u16BitMax / kmhMultiple)
}

func (c Car) newTopSpeed(newSpeed float64) {
	c.topSpeedKMH = newSpeed
}

func newerTopSpeed(c Car, speed float64) Car {
	c.topSpeedKMH = speed
	return c
}

func main() {
	aCar := Car{
		gasPedal:      65000,
		brakePedal:    0,
		steeringWheel: 12561,
		topSpeedKMH:   225.0,
	}

	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh(), "km/h")
	fmt.Println(aCar.milesH(), "miles/h")
	aCar = newerTopSpeed(aCar, 500)
	fmt.Println(aCar.kmh(), "km/h")
	fmt.Println(aCar.milesH(), "miles/h")
}
