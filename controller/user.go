package controller

import (
	"AccountApp/config"
	"AccountApp/users"
	"fmt"
	"strings"
)

// Contoller Register
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

// Controller Login
func Login() (bool, int, string) {
	var phoneNumber, password string

	fmt.Print("Phone number: ")
	fmt.Scanln(&phoneNumber)

	fmt.Print("Password: ")
	fmt.Scanln(&password)

	userModel := &users.UsersModel{DB: config.DB}
	userID, err := userModel.Login(phoneNumber, password)
	if err != nil {
		fmt.Println("Failed to login:", err)
		return false, 0, "" // Mengembalikan nomor telepon sebagai string kosong ("")
	}

	fmt.Println("Login successful!")
	return true, userID, phoneNumber // Mengembalikan nomor telepon yang berhasil login
}

// Controller Read Account
func ViewProfile(phoneNumber string, loggedInUserID int) {
	userModel := &users.UsersModel{DB: config.DB}
	user, err := userModel.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		fmt.Println("Failed to get user:", err)
		return
	}

	if user.ID != loggedInUserID {
		fmt.Println("Error: You can only view your own profile.")
		return
	}

	fmt.Println("Profile:")
	fmt.Println("Phone number:", user.PhoneNumber)
	fmt.Println("Name:", user.Name)
	fmt.Println("Gender:", user.Gender)
	fmt.Println("Tanggal lahir:", user.TanggalLahir)
	fmt.Println("Balance:", user.Balance)
}

// Controller Update Account
func UpdateAccount(phoneNumber string, userID int) {
	var name, gender, tanggalLahir string

	fmt.Print("New Name: ")
	fmt.Scanln(&name)

	fmt.Print("New Gender (M/F): ")
	fmt.Scanln(&gender)
	gender = strings.ToUpper(gender)
	if gender != "M" && gender != "F" {
		fmt.Println("Invalid gender")
		return
	}

	fmt.Print("New Tanggal lahir (YYYY-MM-DD): ")
	fmt.Scanln(&tanggalLahir)

	userModel := &users.UsersModel{DB: config.DB}
	err := userModel.UpdateAccount(userID, name, gender, tanggalLahir)
	if err != nil {
		fmt.Println("Failed to update profile:", err)
		return
	}

	fmt.Println("Profile updated successfully!")
}

// Controller Delete Account
func DeleteAccount(phoneNumber string, userID int) {
	userModel := &users.UsersModel{DB: config.DB}
	err := userModel.DeleteAccount(userID)
	if err != nil {
		fmt.Println("Failed to delete account:", err)
		return
	}

	fmt.Println("Account deleted successfully!")
}

// Controller ViewOtherUserProfile
