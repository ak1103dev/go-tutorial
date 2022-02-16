package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets uint

	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}