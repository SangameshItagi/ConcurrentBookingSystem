package models

type Ticket struct {
	TicketId     uint32 `gorm:"primaryKey" json:"ticketId"`
	SeatReferId  uint32 `json:"seatReferId"`
	EmailReferId string `json:"emailReferId"`
	ShowReferId  uint32 `json:"showReferId"`

	Seat Seat `json:"-"  gorm:"foreignKey:SeatReferId;References:SeatId"`
	User User `json:"-" gorm:"foreignKey:EmailReferId;References:EmailId"`
	Show Show `json:"-" gorm:"foreignKey:ShowReferId;References:ShowId"`
}
