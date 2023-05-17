package controller

import (
	"AccountApp/config"
	"AccountApp/transfer"
	"AccountApp/users"
	"fmt"
)

func TransferAccount(userID int) {
	usersModel := users.UsersModel{DB: config.DB}

	// Get the logged-in user based on the provided userID
	loggedInUser, err := usersModel.GetUserByID(userID)
	if err != nil {
		fmt.Println("", err)
		return
	}

	var SenderPhoneNumber string
	var ReceiverPhoneNumber string
	var Amount int

	fmt.Println("Enter Sender Phone Number:")
	fmt.Scan(&SenderPhoneNumber)

	fmt.Println("Enter Receiver Phone Number:")
	fmt.Scan(&ReceiverPhoneNumber)

	fmt.Println("Enter Amount To Transfer:")
	fmt.Scan(&Amount)

	if Amount <= 0 {
		fmt.Println("saldo tidak boleh Kosong.")
		return
	}

	// Check if the Sender Phone Number matches the logged-in user's Phone Number
	if SenderPhoneNumber != loggedInUser.PhoneNumber {
		fmt.Println("Nomer Pengirim Tidak Sesuai")
		return
	}

	transferModel := transfer.TransferModel{DB: config.DB}
	err = transferModel.TransferAccount(SenderPhoneNumber, ReceiverPhoneNumber, Amount)
	if err != nil {
		fmt.Println("Failed to transfer:", err)
		return
	}

	fmt.Println("Transfer Successful")
}

func PrintTransferHistory(userID int) {
	tm := transfer.TransferModel{DB: config.DB}
	transferHistory, err := tm.GetTransferHistory(userID)
	if err != nil {
		fmt.Println("Failed to retrieve transfer history:", err)
		return
	}

	usersModel := users.UsersModel{DB: config.DB}
	for _, transfer := range transferHistory {
		sender, err := usersModel.GetUserByID(transfer.Sender_ID)
		if err != nil {
			fmt.Println("Failed to get sender user:", err)
			return
		}

		receiver, err := usersModel.GetUserByID(transfer.Receiver_ID)
		if err != nil {
			fmt.Println("Failed to get receiver user:", err)
			return
		}

		fmt.Printf("Transfer ID: %d Sender: %s Receiver: %s Amount: %d Tanggal: %s\n",
			transfer.ID, sender.Name, receiver.Name, transfer.Amount, transfer.Tanggal.Format("2006-01-02 15:04:05"))
	}
}
