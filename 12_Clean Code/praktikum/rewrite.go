package main

type vehicle struct {
	wheel        int
	speedPerHour int
}

type car struct {
	vehicle
}

// function to add speed with parameter int
func (m *car) addSpeed(newSpeed int) {
	m.speedPerHour += newSpeed
}

// function when car is running
func (m *car) run() {
	m.addSpeed(10)
}

func main() {
	fastCar := car{}
	fastCar.run()
	fastCar.run()
	fastCar.run()

	slowCar := car{}
	slowCar.run()

}
