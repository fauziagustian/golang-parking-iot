package attendance

import (
	"errors"
	"math/rand"
	"time"
)

type Attendance struct {
	AttSlot  map[string]string
	Capacity int
}

func (cap *Attendance) GenerateSlotAttend(slotno int) {
	cap.AttSlot = make(map[string]string, slotno)
	cap.Capacity = slotno
}

func (a *Attendance) AttendCar(PlateNumber string) (string, error) {
	attendNo := GenerateAttendNumber(3)
	a.AttSlot[attendNo] = PlateNumber
	return attendNo, nil
}

func (p *Attendance) OutAttend(idAttend string) (string, error) {
	if platNo, found := p.AttSlot[idAttend]; found {
		delete(p.AttSlot, idAttend)
		return platNo, nil
	}
	return "", errors.New("unrecognized Attend")
}

func GenerateAttendNumber(number int) string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("1234567890")

	b := make([]rune, number)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
