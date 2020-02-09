package creational

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufactor := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufactor.SetBuilder(carBuilder)
	manufactor.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on car must be 5 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Wheels on car must be 4 and they were %d\n", car.Wheels)
	}

	bikeBuilder := &BikeBuilder{}
	manufactor.SetBuilder(bikeBuilder)
	manufactor.Construct()

	bike := bikeBuilder.GetVehicle()

	if bike.Wheels != 2 {
		t.Errorf("Wheels on bike must be 2 and they were %d\n", bike.Wheels)
	}

	if bike.Seats != 2 {
		t.Errorf("Seats on bike must be 2 and they were %d\n", bike.Wheels)
	}

	if bike.Structure != "Bike" {
		t.Errorf("Wheels on bike must be 4 and they were %d\n", bike.Wheels)
	}
}
