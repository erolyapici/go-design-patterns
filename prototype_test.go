package main

import "testing"

func TestCarCache_GetClone(t *testing.T) {
	carCahe := getCarCloner()
	if carCahe == nil {
		t.Fatal("Cache was nil")
	}

	item, err := carCahe.GetClone(SUV)
	if err != nil {
		t.Error(err)
	}
	if item == suvPrototype {
		t.Error("item cannot be equal to the suv prototype")
	}

	car, ok := item.(*Car)
	if !ok {
		t.Fatal("Type assertion for car couldn't be done successfully")
	}
	car.Color = "Red"

	cloneItem, err := carCahe.GetClone(SUV)
	if err != nil {
		t.Fatal(err)
	}

	car2, ok := cloneItem.(*Car)
	if !ok {
		t.Fatal("Type assertion for car couldn't be done successfully")
	}
	if car.Color == car2.Color {
		t.Error("Color must be different")
	}
	if car == car2 {
		t.Error("Car cannot be equal car 2")
	}

}
