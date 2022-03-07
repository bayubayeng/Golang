package main

import (
	"fmt"
	"sync"
	"ticketBooking/helper"
	"time"
)

const conferenceTicket int = 50

var remeaningTicket uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName  string
	email     string
	ticket    uint
}

var wd = sync.WaitGroup{}

func main() {

	greetUser()

	firstName, lastName, email, ticket := getUserInput()
	//mevalidasi input user
	isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, ticket, remeaningTicket)

	//seleksi
	if isValidName && isValidEmail && isValidTicket {
		//bookticket
		bookTicket(firstName, lastName, email, ticket)
		wd.Add(1)
		go sendEmail(ticket, firstName, lastName)

		//menulis nama pertama yang memesan
		firstNames := getFirstNames()
		fmt.Println("data nama pertama yang melakukan pemesanan", firstNames)

		if remeaningTicket == 0 {
			//end program
			fmt.Println("mohon maaf ticket hari ini sudah habis silahkan datang  dihari lain")
			//break
		}

	} else {
		if !isValidName {
			fmt.Println("nama depan atau nama belakang yang anda masukan terlalu pendek")
		}
		if !isValidTicket {
			fmt.Println("input ticket yang anda masukan tidak valid")
		}
		if !isValidEmail {
			fmt.Println("email yangn anda masukan tidak ada @")
		}

		//continue
	}
	wd.Wait()
}

func greetUser() {

	fmt.Println("====selamat datang di aplikasi pemesanan hotel====")
	fmt.Printf("hari ini kami menjual %v ticket dan ticket yang tersisa adalah %v \n", conferenceTicket, remeaningTicket)

}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var ticket uint

	//input dari user
	fmt.Print("Nama awal = ")
	fmt.Scanln(&firstName)
	fmt.Print("Nama akhir = ")
	fmt.Scanln(&lastName)
	fmt.Print("email = ")
	fmt.Scanln(&email)
	fmt.Print("ticket yang dipesan = ")
	fmt.Scanln(&ticket)

	return firstName, lastName, email, ticket
}

func bookTicket(firstName string, lastName string, email string, ticket uint) {
	remeaningTicket -= ticket

	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		ticket:    ticket,
	}

	bookings = append(bookings, userData)

	fmt.Printf("daftar pemesan %v\n", bookings)

	fmt.Printf("pesanan ticket sebanyak %d suksess \n sisa ticket yang tersisa adalah %d \n", ticket, remeaningTicket)
	fmt.Printf("bukti pemesanan akan dikirimkan ke email %v\n", email)

}

func sendEmail(ticket uint, firstName string, lastName string) {
	time.Sleep(8 * time.Second)
	user := fmt.Sprintf("%v tiket untuk nama %v %v", ticket, firstName, lastName)
	fmt.Println("====================")
	fmt.Println("email konfirmasi berhasil dikirim ")
	fmt.Println(user)
	fmt.Println("====================")
	wd.Done()
}
