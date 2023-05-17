package transfer

import (
	"AccountApp/users"
	"database/sql"
	"fmt"
)

type TransferModel struct {
	DB *sql.DB
}

func (tm *TransferModel) TransferAccount(SenderPhoneNumber, ReceiverPhoneNumber string, Amount int) error {
	usersModel := users.UsersModel{DB: tm.DB}
	Sender, err := usersModel.GetUserByPhoneNumber(SenderPhoneNumber)
	if err != nil {
		return err
	}

	Receiver, err := usersModel.GetUserByPhoneNumber(ReceiverPhoneNumber)
	if err != nil {
		return err
	}

	if Sender.Balance < Amount {
		return fmt.Errorf("Saldo Tidak Cukup")
	}

	t, err := tm.DB.Begin()
	if err != nil {
		return err
	}
	defer t.Rollback()

	_, err = t.Exec("UPDATE Users SET Balance = Balance - ? WHERE phone_number = ?", Amount, SenderPhoneNumber)
	if err != nil {
		return err
	}

	_, err = t.Exec("UPDATE Users SET Balance = Balance + ? WHERE phone_number = ?", Amount, ReceiverPhoneNumber)
	if err != nil {
		return err
	}

	_, err = t.Exec("INSERT INTO transfer (sender_id, receiver_id, amount, tanggal) VALUES (?,?,?,NOW())", Sender.ID, Receiver.ID, Amount)
	if err != nil {
		return err
	}

	err = t.Commit()
	if err != nil {
		return err
	}

	return nil
}
