package topup

import (
	"database/sql"
	"errors"
)

type TopupModel struct {
	DB *sql.DB
}

func (tm *TopupModel) TopUpAccount(phoneNumber string, amount int) error {
	// Get the user ID based on the phone number
	var userID int
	err := tm.DB.QueryRow("SELECT id FROM users WHERE phone_number = ?", phoneNumber).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("user not found")
		}
		return err
	}

	// Insert top-up record
	_, err = tm.DB.Exec("INSERT INTO topup (user_id, amount, tanggal) VALUES (?, ?, NOW())", userID, amount)
	if err != nil {
		return err
	}

	// Update user balance
	_, err = tm.DB.Exec("UPDATE users SET balance = balance + ? WHERE id = ?", amount, userID)
	if err != nil {
		return err
	}

	return nil
}
