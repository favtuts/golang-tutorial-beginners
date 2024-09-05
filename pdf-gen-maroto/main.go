package main

type Company struct {
	Name         string
	Address      string
	LogoLocation string
}

type Ticket struct {
	ID                 int
	ShowName           string
	ShowTime           string
	Language           string
	ShowVenue          string
	SeatNumber         string
	Cost               float64
	Screen             string
	TicketCount        int
	ShowPosterLocation string
}

func main() {
	c := Company{
		Name:         "ShowBees Ticketing",
		Address:      "1234 Main St, City, State 12345",
		LogoLocation: "./logo.png",
	}

	t := Ticket{
		ID:                 1,
		ShowName:           "Planet of the Gophers: The War Begins",
		ShowTime:           "Sat 01/01/2022 7:00 PM",
		Language:           "English",
		ShowVenue:          "Gophedorium",
		SeatNumber:         "Platinum - A1, A2",
		Cost:               620.00,
		Screen:             "Screen 1",
		TicketCount:        2,
		ShowPosterLocation: "./poster.png",
	}
}
