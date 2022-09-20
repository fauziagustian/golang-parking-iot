package parking

import (
	"errors"
	"math/rand"
	"time"
)

// Parking is main component, for data structure a parking lot
type Parking struct {
	Slot     map[string]string
	Capacity int
}

func (cap *Parking) GenerateSlot(slotno int) {
	cap.Slot = make(map[string]string, slotno)
	cap.Capacity = slotno
}

func (p *Parking) ParkingCar(number string) (string, error) {

	if p.Capacity == len(p.Slot) {
		return "", errors.New("no available position")
	}

	if p.checkPlatNoExist(number) {
		return "", errors.New("car plat no already exist")
	}

	ticketNo := generateTicketNo(5)
	p.Slot[ticketNo] = number

	return ticketNo, nil
}

func (p *Parking) checkPlatNoExist(platNo string) bool {
	for _, value := range p.Slot {
		if value == platNo {
			return true
		}
	}

	return false
}

func (p *Parking) CheckCarByNumber(number string) string {
	var desc string
	desc = "Plat Number does not exist"
	for _, value := range p.Slot {
		if value == number {
			desc = "Plat Number Exist"
		}
	}
	return desc
}

func (p *Parking) OutParkCar(ticketNo string) (string, error) {
	if platNo, found := p.Slot[ticketNo]; found {
		delete(p.Slot, ticketNo)
		return platNo, nil
	}
	return "", errors.New("unrecognized parking ticket")
}

func generateTicketNo(n int) string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
