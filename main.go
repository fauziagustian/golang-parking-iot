package main

import (
	"exercise-go-parking-Iot/attendance"
	"exercise-go-parking-Iot/car"
	"exercise-go-parking-Iot/parking"
	"fmt"
)

func main() {
	Parking := parking.Parking{}
	Parking.GenerateSlot(2)
	Attendance := attendance.Attendance{}
	Attendance.GenerateSlotAttend(2)

	carParking := map[string]string{
		"1": "f 5667 rg",
		"2": "b 3434 fec",
	}
	carAttend := map[string]string{
		"1": "1",
		"2": "2",
	}

	fmt.Println(carParking)
	fmt.Println(carAttend)

	cars := car.Car{Number: "F 1234 ER"}
	_, err := Parking.ParkingCar(cars.Number)
	if err != nil {
		println(err)
	} else {
		_, err := Attendance.AttendCar(cars.Number)
		if err != nil {
			println(err)
		}
	}

	fmt.Println("===================STEP 1===================")
	fmt.Println(Parking.Slot)
	fmt.Println(Parking.CheckCarByNumber("F 1235 ER"))
	fmt.Println(Parking.CheckCarByNumber("F 1234 ER"))
	fmt.Println(Attendance.AttSlot)

	cars2 := car.Car{Number: "F 4321 RE"}
	ticketNoCar2, err := Parking.ParkingCar(cars2.Number)
	if err != nil {
		println(err)
		return
	}

	idAttend, err := Attendance.AttendCar(cars2.Number)
	if err != nil {
		println(err)
	}

	fmt.Println("===================STEP 2===================")
	fmt.Println("Check Slot After Add more than 1 : \n", Parking.Slot)
	fmt.Println(Parking.CheckCarByNumber("F 3532 ER"))
	fmt.Println(Parking.CheckCarByNumber("F 4321 RE"))
	fmt.Println(Attendance.AttSlot)
	//fmt.Println(ticketNoCar2)

	//use it test if absent id not same
	//idAttend = "2432"
	_, err = Attendance.OutAttend(idAttend)
	if err != nil {
		println(err)
	} else {
		_, err := Parking.OutParkCar(ticketNoCar2)
		if err != nil {
			println(err)
		}
	}

	fmt.Println("===================STEP 3===================")
	//fmt.Println(outParkCars2)
	fmt.Println(Parking.Slot)
	fmt.Println(Attendance.AttSlot)

}
