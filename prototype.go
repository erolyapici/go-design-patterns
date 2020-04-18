package main

import (
	"errors"
	"fmt"
)

type CarCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	SUV   = 1
	HB    = 2
	SEDAN = 3
)

func getCarCloner() CarCloner {
	return new(CarCache)
}

type CarCache struct {
}

func (s *CarCache) GetClone(a int) (ItemInfoGetter, error) {
	switch a {
	case SUV:
		car := *suvPrototype
		return &car, nil
	case HB:
		car := *hbPrototype
		return &car, nil
	case SEDAN:
		car := *sedanPrototype
		return &car, nil
	default:
		return nil, errors.New("Not implemented yet")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}
type CarType byte
type Car struct {
	Price float32
	Type  CarType
	Color string
}

func (s *Car) GetInfo() string {
	return fmt.Sprintf("Car type=%d color=%s Price=%f", s.Type, s.Color, s.Price)
}

var suvPrototype *Car = &Car{
	Price: 1,
	Type:  SUV,
	Color: "White",
}
var hbPrototype *Car = &Car{
	Price: 1,
	Type:  HB,
	Color: "Red",
}
var sedanPrototype *Car = &Car{
	Price: 1,
	Type:  SEDAN,
	Color: "Black",
}

func (i *Car) GetPrice() float32 {
	return i.Price
}
