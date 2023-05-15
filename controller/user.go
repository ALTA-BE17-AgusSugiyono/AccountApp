package controller

import (
	"AccountApp/config"
	"AccountApp/users"
	"fmt"
	"strings"
)

func Register() {
	var phoneNumber, password, name, gender, tanggalLahir string

	fmt.Print("Phone number: ")
	fmt.Scanln(&phoneNumber)

	fmt.Print("Password: ")
	fmt.Scanln(&password)

	fmt.Print("Name: ")
	fmt.Scanln(&name)

	fmt.Print("Gender (M/F): ")
	fmt.Scanln(&gender)
	gender = strings.ToUpper(gender)
	if gender != "M" && gender != "F" {
		fmt.Println("Invalid gender")
		return
	}

	fmt.Print("Tanggal lahir (YYYY-MM-DD): ")
	fmt.Scanln(&tanggalLahir)

	// buat instance dari UsersModel dan inisialisasi dengan DB yang telah diinisialisasi pada config.Init()
	userModel := &users.UsersModel{DB: config.DB}

	err := userModel.Register(phoneNumber, password, name, gender, tanggalLahir)
	if err != nil {
		fmt.Println("Failed to register user:", err)
		return
	}

	fmt.Println("User registered successfully!")
}
