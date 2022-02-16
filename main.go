package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have tottal of", conferenceTickets, "tickets and", remainingTickets, "are still avaliable.")
	fmt.Println("Get your tickets here to attend")
}