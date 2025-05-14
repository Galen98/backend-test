package models

type Konsumsi struct {
	Name string `json:"name"`
}

type Booking struct {
	BookingDate     string     `json:"bookingDate"`
	OfficeName      string     `json:"officeName"`
	StartTime       string     `json:"startTime"`
	EndTime         string     `json:"endTime"`
	ListConsumption []Konsumsi `json:"listConsumption"`
	Participants    int        `json:"participants"`
	RoomName        string     `json:"roomName"`
	ID              string     `json:"id"`
}
