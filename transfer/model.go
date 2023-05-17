package transfer

import (
	"AccountApp/users"
	"database/sql"
	"fmt"
	"time"
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

func (tm *TransferModel) GetTransferHistory(userID int) ([]Transfer, error) {
	rows, err := tm.DB.Query("SELECT id, sender_id, receiver_id, amount, tanggal FROM transfer WHERE sender_id = ? OR receiver_id = ?", userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	history := []Transfer{}
	for rows.Next() {
		var t Transfer
		var tanggalString string
		err := rows.Scan(&t.ID, &t.Sender_ID, &t.Receiver_ID, &t.Amount, &tanggalString)
		if err != nil {
			return nil, err
		}
		t.Tanggal, err = time.Parse("2006-01-02 15:04:05", tanggalString)
		if err != nil {
			return nil, err
		}
		history = append(history, t)
	}

	return history, nil
}
