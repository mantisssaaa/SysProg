package main

import (
	"fmt"
	"time"
)

type Room struct {
	number   string
	roomType string
	price    float64
	isBooked bool
}

type Hotel struct {
	name  string
	rooms []Room
}

type Reservation struct {
	hotel     string
	room      string
	guestName string
	startDate time.Time
	endDate   time.Time
	totalCost float64
}

func NewHotel(name string) *Hotel {
	return &Hotel{
		name:  name,
		rooms: make([]Room, 0),
	}
}

func (h *Hotel) AddRoom(number string, roomType string, price float64) {
	room := Room{
		number:   number,
		roomType: roomType,
		price:    price,
		isBooked: false,
	}
	h.rooms = append(h.rooms, room)
}

func (h *Hotel) CheckAvailability(roomType string, startDate time.Time, endDate time.Time) []Room {
	var availableRooms []Room
	for _, room := range h.rooms {
		if room.roomType == roomType && !room.isBooked {
			availableRooms = append(availableRooms, room)
		}
	}
	return availableRooms
}

func (h *Hotel) MakeReservation(roomNumber string, guestName string, startDate time.Time, endDate time.Time) *Reservation {
	for i, room := range h.rooms {
		if room.number == roomNumber && !room.isBooked {
			days := endDate.Sub(startDate).Hours() / 24
			totalCost := room.price * days

			h.rooms[i].isBooked = true

			return &Reservation{
				hotel:     h.name,
				room:      roomNumber,
				guestName: guestName,
				startDate: startDate,
				endDate:   endDate,
				totalCost: totalCost,
			}
		}
	}
	return nil
}

func main() {
	hotel := NewHotel("Отель")

	hotel.AddRoom("101", "Одноместная", 3000)
	hotel.AddRoom("102", "Двуместная", 4000)
	hotel.AddRoom("103", "Одноместная", 3000)
	hotel.AddRoom("201", "Трехместная", 5000)
	available := hotel.CheckAvailability("Одноместная", time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC), time.Date(2025, 1, 20, 0, 0, 0, 0, time.UTC))
	fmt.Printf("Доступные одноместные номера: %v\n", available)

	reservation := hotel.MakeReservation("101", "Иван Сэргеевич", time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC), time.Date(2025, 1, 20, 0, 0, 0, 0, time.UTC))
	if reservation != nil {
		fmt.Printf("\nБронирование успешно!\n")
		fmt.Printf("Отель: %v\n", reservation.hotel)
		fmt.Printf("Комната: %v\n", reservation.room)
		fmt.Printf("Гость: %v\n", reservation.guestName)
		fmt.Printf("Даты: %v - %v\n", reservation.startDate.Format("2006-01-02"), reservation.endDate.Format("2006-01-02"))
		fmt.Printf("Общая стоимость: %v руб.\n", reservation.totalCost)
		availableAfter := hotel.CheckAvailability("Одноместная", time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC), time.Date(2025, 1, 20, 0, 0, 0, 0, time.UTC))
		fmt.Printf("\nДоступные комнаты Standard после бронирования: %v\n", availableAfter)
	} else {
		fmt.Println("Не удалось забронировать комнату")
	}
}
