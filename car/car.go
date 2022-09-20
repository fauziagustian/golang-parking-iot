package car

type Car struct {
	Number   string
	TicketNo string
}

func (this *Car) IsEqual(cr Car) bool {
	return (this.Number == cr.Number && this.TicketNo == cr.TicketNo)
}
