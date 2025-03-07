package vehicles

import "fmt"

type Vehicle interface {
	Distance() float64
}

type Car struct {
	Time int
}

type Motorcycle struct {
	Time int
}

type Truck struct {
	Time int
}

type Sarasa struct {
	Time int
}

const (
	CarVechicle        = "Car"
	MotorcycleVechicle = "Motorcycle"
	TruckVehicle       = "Truck"
	SarasaVehicle      = "Sarasa"
)

func New(vehicle string, time int) (Vehicle, error) {
	switch vehicle {
	case CarVechicle:
		return &Car{Time: time}, nil
	case MotorcycleVechicle:
		return &Motorcycle{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	case SarasaVehicle:
		return &Sarasa{Time: time}, nil
	}
	return nil, fmt.Errorf("vehicle '%s' not exists", vehicle)
}

func (c *Car) Distance() float64 {
	return 100 * (float64(c.Time) / 60)
}

func (c *Motorcycle) Distance() float64 {
	return 120 * (float64(c.Time) / 60)
}

func (c *Truck) Distance() float64 {
	return 70 * (float64(c.Time) / 60)
}

func (c *Sarasa) Distance() float64 {
	return 300000 * (float64(c.Time) / 60)
}
