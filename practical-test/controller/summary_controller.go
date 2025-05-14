package controller

import (
	"encoding/json"
	"io/ioutil"
	"logical-test/practical-test/models"
	"net/http"
)

type DetailItem struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Total int    `json:"total"`
}

type SummaryItem struct {
	RoomName     string       `json:"roomName"`
	Total        int          `json:"total"`
	Participants int          `json:"participants"`
	Detail       []DetailItem `json:"details"`
}

func GetSummary(w http.ResponseWriter, r *http.Request) {
	bookingRes, err := http.Get("https://66876cc30bc7155dc017a662.mockapi.io/api/dummy-data/bookingList")
	if err != nil {
		http.Error(w, "error fetching booking data", http.StatusInternalServerError)
		return
	}
	defer bookingRes.Body.Close()

	body, _ := ioutil.ReadAll(bookingRes.Body)
	var bookings []models.Booking
	json.Unmarshal(body, &bookings)

	masterRes, err := http.Get("https://6686cb5583c983911b03a7f3.mockapi.io/api/dummy-data/masterJenisKonsumsi")
	if err != nil {
		http.Error(w, "error fetching master konsumsi", http.StatusInternalServerError)
		return
	}
	defer masterRes.Body.Close()

	masterBody, _ := ioutil.ReadAll(masterRes.Body)
	var masterList []models.MasterKonsumsi
	json.Unmarshal(masterBody, &masterList)

	priceMap := make(map[string]int)
	for _, item := range masterList {
		priceMap[item.Name] = item.MaxPrice
	}
	roomMap := make(map[string]*SummaryItem)

	for _, booking := range bookings {
		room := booking.RoomName

		if _, exists := roomMap[room]; !exists {
			roomMap[room] = &SummaryItem{
				RoomName:     room,
				Participants: 0,
				Detail:       []DetailItem{},
				Total:        0,
			}
		}

		roomSummary := roomMap[room]
		roomSummary.Participants += booking.Participants

		for _, konsumsi := range booking.ListConsumption {
			price := priceMap[konsumsi.Name]
			found := false

			for i, d := range roomSummary.Detail {
				if d.Name == konsumsi.Name {
					roomSummary.Detail[i].Total += price * booking.Participants
					found = true
					break
				}
			}

			if !found {
				roomSummary.Detail = append(roomSummary.Detail, DetailItem{
					Name:  konsumsi.Name,
					Price: price,
					Total: price * booking.Participants,
				})
			}

			roomSummary.Total += price * booking.Participants
		}
	}

	var summarys []SummaryItem
	for _, summary := range roomMap {
		summarys = append(summarys, *summary)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summarys)
}
